package document

import (
	"context"
	"digitalsignature/internal/app/model"
	"digitalsignature/internal/app/repository"
	"fmt"
	"log"
	"math/big"
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

type DocumentService interface {
	VerifyDocument(phone string, digest [32]byte, docNum *big.Int) (bool, error)
	SaveSignaturalDocument(phone string, signatural []byte) (Event, error)
	GetDocumentByPublickey(publickey string) ([]model.Document, error)
}

type document struct {
	accountSrv   AccountSrv
	documentRepo *repository.DocumentRepository
	client       *ethclient.Client
	instance     *Document
	log          *zap.Logger
	address      string
}

type Event struct {
	Publickey string `json:"publickey"`
	Numdoc    uint   `json:"numdoc"`
	Signature []byte `json:"signature"`
}

func NewDocumentService(
	client *ethclient.Client,
	userAddress common.Address,
	accountSrv AccountSrv,
	documentRepo *repository.DocumentRepository,
	address string,
	log *zap.Logger) DocumentService {

	doc, err := NewDocument(userAddress, client)
	if err != nil {
		log.Sugar().Error(err)
	}

	return &document{
		accountSrv:   accountSrv,
		documentRepo: documentRepo,
		client:       client,
		instance:     doc,
		log:          log,
		address:      address,
	}
}

func (d *document) GetDocumentByPublickey(publickey string) ([]model.Document, error) {
	docs, err := d.documentRepo.AllByOwner(publickey)
	if err != nil {
		log.Fatal(err)
	}
	return docs, nil
}

// Verify document by using phone, digest, DocID
// phone using for get public key
func (d *document) VerifyDocument(phone string, digest [32]byte, docNum *big.Int) (bool, error) {
	isCorrect, err := d.instance.VerifyDoc(&bind.CallOpts{}, phone, digest, docNum)
	return isCorrect, err
}

func (d *document) SaveSignaturalDocument(phone string, signatural []byte) (Event, error) {
	fmt.Println("phone:", phone)
	fmt.Println("signatural:", signatural)
	account, ks, err := d.accountSrv.GetAminAccount()
	if err != nil {
		d.log.Sugar().Error(err)
	}
	opts := d.accountSrv.BindTransactionOption(*account, "123456", ks, d.client)

	txn, err := d.instance.SaveDoc(opts, phone, signatural)
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

	address := common.HexToAddress(d.address)
	d.log.Info("contract address: " + address.String())
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   currentBlock,
		Addresses: []common.Address{
			address,
		},
	}
	logs, err := d.client.FilterLogs(context.Background(), query)
	if err != nil {
		d.log.Sugar().Error(err)
		return Event{}, err
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(DocumentABI)))
	if err != nil {
		d.log.Sugar().Error(err)
		return Event{}, err
	}
	events := []Event{}
	for _, vLog := range logs {
		event := Event{}
		err = contractAbi.UnpackIntoInterface(&event, "IndexDocument", vLog.Data)
		if err != nil {
			d.log.Sugar().Error(err)
		}
		events = append(events, event)
	}
	if len(events) > 0 {
		fmt.Println(events)
		return events[len(events)-1], nil
	}
	return Event{}, nil
}
