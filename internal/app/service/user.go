package service

import (
	"digitalsignature/internal/app/model"
	"digitalsignature/internal/app/repository"
	"digitalsignature/internal/app/request"
	"digitalsignature/internal/app/service/document"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type AccountSrv interface {
	GetAminAccount() (*accounts.Account, *keystore.KeyStore, error)
	BindTransactionOption(account accounts.Account, password string, ks *keystore.KeyStore, client *ethclient.Client) *bind.TransactOpts
}
type UserService struct {
	ethclient        *ethclient.Client
	userRepo         *repository.UserRepo
	logger           *zap.Logger
	accountSrv       AccountSrv
	documentContract string
}

func NewUserService(client *ethclient.Client, userRepo *repository.UserRepo, accountSrv AccountSrv, docContract string, log *zap.Logger) *UserService {
	return &UserService{
		ethclient:        client,
		userRepo:         userRepo,
		logger:           log,
		accountSrv:       accountSrv,
		documentContract: docContract,
	}
}
func (s *UserService) Create(c *gin.Context, userInfo request.UserInfo) error {
	// Validate user info including gmail, phone, id card

	// Insert user info to offchain
	user := model.User{
		ID:          uuid.New(),
		Name:        userInfo.Name,
		PublicKey:   userInfo.PublicKey,
		CardID:      userInfo.CardID,
		Phone:       userInfo.Phone,
		Gmail:       userInfo.Email,
		DateOfBirth: userInfo.DateOfBirth,
		CreateAt:    time.Now(),
	}
	err := s.userRepo.Create(user)
	if err != nil {
		s.logger.Error("create new account error", zap.String("error", err.Error()))
	}
	// make transaction with admin account
	adminAccount, ks, err := s.accountSrv.GetAminAccount()
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(s.documentContract)
	// store hash user info to blockchain
	documentIntance, err := document.NewDocument(contractAddress, s.ethclient)
	if err != nil {
		log.Fatal(err)
	}
	tnxOption := s.accountSrv.BindTransactionOption(*adminAccount, "123456", ks, s.ethclient)
	txn, err := documentIntance.StoreUser(tnxOption, user.ID.String(), user.Name, user.CardID, user.DateOfBirth, user.Phone, user.Gmail, user.PublicKey)
	if err != nil {
		log.Fatal(err)
	}
	var wg1 sync.WaitGroup
	wg1.Add(1)
	go func() {
		for {
			receipt, err := s.ethclient.TransactionReceipt(c, txn.Hash())
			if err != nil {
				log.Fatal(err)
			}
			if receipt.Status == 1 || receipt.Status == 0 {
				wg1.Done()
				if receipt.Status == 1 {
					s.logger.Info("Transaction Success")
				}
				if receipt.Status == 0 {
					s.logger.Info("Transaction Failt")
				}
				break
			}
		}
	}()
	wg1.Wait()
	s.logger.Info("transaction hash", zap.String("hash", txn.Hash().String()))
	time.Sleep(3 * time.Second)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			receipt, err := s.ethclient.TransactionReceipt(c, txn.Hash())
			if err != nil {
				fmt.Println("transaction not found")
				log.Fatal(err)
			}
			if receipt.Status == 1 || receipt.Status == 0 {
				wg.Done()
				if receipt.Status == 1 {
					s.logger.Info("Transaction Success")
				}
				if receipt.Status == 0 {
					s.logger.Info("Transaction Failt")
				}
				break
			}
		}
	}()
	wg.Wait()
	s.logger.Info(txn.Hash().String())
	s.logger.Info("created user", zap.String("publickey", userInfo.PublicKey))
	return nil
}

func (s *UserService) GetHashUserInfo(c *gin.Context, pubkey string) error {
	contractAddress := common.HexToAddress(s.documentContract)
	// store hash user info to blockchain
	documentIntance, err := document.NewDocument(contractAddress, s.ethclient)
	if err != nil {
		log.Fatal(err)
	}
	hash, err := documentIntance.GetHashUserInfo(&bind.CallOpts{}, pubkey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hash)
	return nil
}
