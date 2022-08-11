package document

import (
	"context"
	"digitalsignature/internal/app/model"
	"digitalsignature/internal/app/repository"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type AccountSrv interface {
	GetAminAccount() (*accounts.Account, *keystore.KeyStore, error)
	BindTransactionOption(
		account accounts.Account,
		password string,
		ks *keystore.KeyStore,
		client *ethclient.Client) *bind.TransactOpts
}

type DocumentService struct {
	accountSrv      AccountSrv
	documentRepo    *repository.DocumentRepo
	userRepo        *repository.UserRepo
	contractAddress common.Address
	client          *ethclient.Client
	log             *zap.Logger
}

type Event struct {
	Numdoc *big.Int
}

func NewDocumentService(
	client *ethclient.Client,
	accountSrv AccountSrv,
	documentRepo *repository.DocumentRepo,
	userRepo *repository.UserRepo,
	log *zap.Logger) *DocumentService {

	contractAddress, ok := os.LookupEnv("CONTRACT_ADDRESS")
	if !ok {
		return nil
	}
	return &DocumentService{
		accountSrv:      accountSrv,
		documentRepo:    documentRepo,
		userRepo:        userRepo,
		contractAddress: common.HexToAddress(contractAddress),
		client:          client,
		log:             log,
	}
}

func (d *DocumentService) GetDocumentByPhone(phone string) ([]model.Document, error) {
	userPublicKey, err := d.userRepo.GetPublickeyByPhone(phone)
	if err != nil {
		d.log.Sugar().Error(err)
		return []model.Document{}, err
	}
	docs, err := d.documentRepo.AllByOwner([]string{userPublicKey})
	if err != nil {
		d.log.Sugar().Error(err)
		return []model.Document{}, err
	}
	return docs, nil
}

func (d *DocumentService) GetDocumentByPublickey(publickeys []string) ([]model.Document, error) {
	docs, err := d.documentRepo.AllByOwner(publickeys)
	if err != nil {
		d.log.Sugar().Error(err)
		return []model.Document{}, err
	}
	return docs, nil
}

func (d *DocumentService) GetSignature(phone string, number *big.Int) ([]byte, error) {
	// get hash user from onchain
	documentIntance, err := NewDocument(d.contractAddress, d.client)
	if err != nil {
		d.log.Sugar().Error(err)
		return []byte{}, err
	}
	signature, _, err := documentIntance.GetSignature(&bind.CallOpts{}, phone, number)
	if err != nil {
		d.log.Sugar().Error(err)
		return []byte{}, err
	}
	return signature, nil
}

// Verify document by using phone, digest, DocID
// phone using for get public key
func (d *DocumentService) VerifyDocument(phone string, digest string, docNum *big.Int) (bool, error) {
	// get hash user from onchain
	documentIntance, err := NewDocument(d.contractAddress, d.client)
	if err != nil {
		d.log.Sugar().Error(err)
		return false, err
	}
	isCorrect, err := documentIntance.VerifyDoc(&bind.CallOpts{}, phone, digest, docNum)
	return isCorrect, err
}

func (d *DocumentService) SaveSignaturalDocument(phone string, signatural []byte) (Event, error) {
	fmt.Println("phone:", phone)
	fmt.Println("signatural:", signatural)
	account, ks, err := d.accountSrv.GetAminAccount()
	if err != nil {
		d.log.Sugar().Error(err)
	}
	opts := d.accountSrv.BindTransactionOption(*account, "123456", ks, d.client)
	// get hash user from onchain
	documentIntance, err := NewDocument(d.contractAddress, d.client)
	if err != nil {
		d.log.Sugar().Error(err)
		return Event{}, err
	}
	txn, err := documentIntance.SaveDoc(opts, phone, signatural)
	if err != nil {
		d.log.Sugar().Error(err)
		return Event{}, err
	}
	var wg1 sync.WaitGroup
	time.Sleep(5 * time.Second)

	wg1.Add(1)
	go func() {
		for {
			receipt, err := d.client.TransactionReceipt(context.Background(), txn.Hash())
			if err != nil {
				log.Println("transaction save document: ", txn.Hash(), ":", err)
				break
			} else {
				if receipt.Status == 1 || receipt.Status == 0 {
					defer wg1.Done()
					if receipt.Status == 1 {
						d.log.Info("Transaction Success")
					}
					if receipt.Status == 0 {
						d.log.Info("Transaction Failt")
					}
					break
				}
			}
			time.Sleep(2 * time.Second)
		}
	}()
	wg1.Wait()
	// check event get document number
	header, err := d.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		d.log.Sugar().Error(err)

	}

	currentBlock := header.Number
	fromBlock := big.NewInt(currentBlock.Int64() - int64(100))

	d.log.Info("from block", zap.Int("block", int(fromBlock.Int64())))

	d.log.Info("current block", zap.Int("block", int(currentBlock.Int64())))
	d.log.Info("contract address: " + d.contractAddress.String())
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   currentBlock,
		Addresses: []common.Address{
			d.contractAddress,
		},
	}
	logs, err := d.client.FilterLogs(context.Background(), query)
	if err != nil {
		d.log.Sugar().Error(err)
		return Event{}, err
	}
	d.log.Info("logs length ", zap.Int("length", len(logs)))

	contractAbi, err := abi.JSON(strings.NewReader(string(DocumentABI)))
	if err != nil {
		d.log.Sugar().Error(err)
		return Event{}, err
	}
	events := []Event{}
	for _, vLog := range logs {
		event := Event{}
		err = contractAbi.UnpackIntoInterface(&event, "Number", vLog.Data)
		if err != nil {
			d.log.Sugar().Error(err)
		}
		events = append(events, event)
	}
	if len(events) > 0 {
		fmt.Println(events)
		// Update doc_num trong database
		num := int(events[len(events)-1].Numdoc.Int64())
		err := d.documentRepo.UpdateDocumentNumber(string(signatural), num)
		if err != nil {
			d.log.Sugar().Error(err)
		}
		d.log.Info("document", zap.Int("number", num))
		return events[len(events)-1], nil
	}
	return Event{}, nil
}
