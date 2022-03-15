package service

import (
	"context"
	"digitalsignature/internal/app/model"
	"digitalsignature/internal/app/repository"
	"digitalsignature/internal/app/request"
	"digitalsignature/internal/app/response"
	"digitalsignature/internal/app/service/document"
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
		Gmail:       userInfo.Gmail,
		DateOfBirth: userInfo.DateOfBirth,
		Password:    userInfo.Password,
		CreateAt:    time.Now(),
	}
	err := s.userRepo.Create(user)
	if err != nil {
		s.logger.Error("create new account error", zap.String("error", err.Error()))
	}
	// make transaction with admin account
	adminAccount, ks, err := s.accountSrv.GetAminAccount()
	if err != nil {
		s.logger.Sugar().Error(err)
	}
	contractAddress := common.HexToAddress(s.documentContract)
	// store hash user info to blockchain
	documentIntance, err := document.NewDocument(contractAddress, s.ethclient)
	if err != nil {
		s.logger.Sugar().Error(err)
	}
	tnxOption := s.accountSrv.BindTransactionOption(*adminAccount, "123456", ks, s.ethclient)
	txn, err := documentIntance.StoreUser(tnxOption, user.ID.String(), user.Name, user.CardID, user.DateOfBirth, user.Phone, user.Gmail, user.PublicKey)
	if err != nil {
		s.logger.Sugar().Error(err)
	}
	var wg1 sync.WaitGroup
	time.Sleep(5 * time.Second)
	wg1.Add(1)
	go func() {
		for {
			receipt, err := s.ethclient.TransactionReceipt(context.Background(), txn.Hash())
			if err != nil {
				log.Println("transaction store user: ", txn.Hash(), ":", err)
			} else {
				if receipt.Status == 1 || receipt.Status == 0 {
					defer wg1.Done()
					if receipt.Status == 1 {
						s.logger.Info("Transaction Success")
					}
					if receipt.Status == 0 {
						s.logger.Info("Transaction Failt")
					}
					break
				}
			}
			time.Sleep(2 * time.Second)
		}
	}()
	wg1.Wait()
	s.logger.Info("created user", zap.String("publickey", userInfo.PublicKey))
	return nil
}

func (s *UserService) GetUserInfo(c *gin.Context, pubkey string) (*response.User, error) {
	contractAddress := common.HexToAddress(s.documentContract)
	// get hash user from onchain
	documentIntance, err := document.NewDocument(contractAddress, s.ethclient)
	if err != nil {
		s.logger.Sugar().Error(err)
		return nil, err
	}
	hash, err := documentIntance.GetHashUserInfo(&bind.CallOpts{}, pubkey)
	if err != nil {
		s.logger.Sugar().Error(err)
		return nil, err
	}
	// get user info from offchain
	user, err := s.userRepo.GetUserByPubkey(pubkey)
	if err != nil {
		s.logger.Sugar().Error(err)
		return nil, err
	}
	return &response.User{
		ID:          user.ID,
		Name:        user.Name,
		PublicKey:   pubkey,
		CardID:      user.CardID,
		Phone:       user.Phone,
		Gmail:       user.Gmail,
		DateOfBirth: user.DateOfBirth,
		Hash:        string(hash[:]),
	}, nil
}

func (s *UserService) VerifyUser(user request.UserInfo) (bool, error) {
	contractAddress := common.HexToAddress(s.documentContract)
	// get hash user from onchain
	documentIntance, err := document.NewDocument(contractAddress, s.ethclient)
	if err != nil {
		s.logger.Sugar().Error(err)
		return false, err
	}
	isVerify, err := documentIntance.VerifyUser(&bind.CallOpts{}, user.Name, user.CardID, user.DateOfBirth, user.Phone, user.Gmail, user.PublicKey)
	return isVerify, err
}

func (s *UserService) Login(login request.Login) (bool, *model.User, error) {
	return s.userRepo.CheckLogin(login.Password, login.Phone)
}
