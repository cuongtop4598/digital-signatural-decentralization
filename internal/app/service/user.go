package service

import (
	"context"
	"digitalsignature/internal/app/errors"
	"digitalsignature/internal/app/model"
	"digitalsignature/internal/app/repository"
	"digitalsignature/internal/app/request"
	"digitalsignature/internal/app/response"
	"digitalsignature/internal/pkg/abi"
	"digitalsignature/internal/pkg/async"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AccountSrv interface {
	GetBindTransactionOptions(client *ethclient.Client) *bind.TransactOpts
}

type UserService struct {
	ethclient        *ethclient.Client
	userRepo         *repository.UserRepo
	logger           *zap.Logger
	accountSrv       AccountSrv
	documentContract string
}

func NewUserService(client *ethclient.Client, userRepo *repository.UserRepo, accountSrv AccountSrv, log *zap.Logger) *UserService {
	contractAddress, _ := os.LookupEnv("CONTRACT_ADDRESS")
	return &UserService{
		ethclient:        client,
		userRepo:         userRepo,
		logger:           log,
		accountSrv:       accountSrv,
		documentContract: contractAddress,
	}
}

func (s *UserService) Register(c *gin.Context, userInfo request.UserInfo) error {
	registerStatus := false
	s.logger.Info("USER_REGISTER", zap.Any("Data", userInfo))
	// TODO: Validate user info including gmail, phone, id card

	// TODO: Check user is exist in database?
	if s.userRepo.IsExist(userInfo.Phone, userInfo.PublicKey) {
		// Insert valid user
		user := model.User{
			Name:        userInfo.Name,
			PublicKey:   userInfo.PublicKey,
			CardID:      userInfo.CardID,
			Phone:       userInfo.Phone,
			Gmail:       userInfo.Gmail,
			DateOfBirth: userInfo.DateOfBirth,
			Password:    userInfo.Password,
			CreateAt:    time.Now(),
		}

		// make transaction with admin account
		contractAddress := common.HexToAddress(s.documentContract)
		// store hash user info to blockchain
		documentIntance, err := abi.NewAbi(contractAddress, s.ethclient)
		if err != nil {
			s.logger.Sugar().Error(err)
		}
		address := common.HexToAddress(user.PublicKey)
		txOption := s.accountSrv.GetBindTransactionOptions(s.ethclient)
		storeUserAsync := async.Exec(func() (*types.Transaction, error) {
			return documentIntance.StoreUser(txOption, user.Name, user.CardID, user.DateOfBirth, user.Phone, user.Gmail, address)
		})
		txn, _ := storeUserAsync.Await()
		if err != nil {
			s.logger.Sugar().Error(err)
		}
		time.Sleep(5 * time.Second)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			for {
				receipt, err := s.ethclient.TransactionReceipt(context.Background(), txn.Hash())
				if err != nil {
					s.logger.Info("TRANSACTION_STORE_USER", zap.Any(txn.Hash().String(), "Is Pending"))
					time.Sleep(5 * time.Second)
				} else {
					wg.Done()
					if receipt.Status == 1 || receipt.Status == 0 {
						if receipt.Status == 1 {
							s.logger.Info("TRANSACTION_STORE_USER_STATUS", zap.Any("Status", "Success"))
							// out, _ := txn.MarshalJSON()
							// s.logger.Info("HASH USER INFO", zap.String("Hash", string(out)))
							registerStatus = true
						}
						if receipt.Status == 0 {
							s.logger.Warn("TRANSACTION_STORE_USER_STATUS STORE USER STATUS", zap.Any("Status", "Failt"))
						}
					}
					break
				}
			}
		}()
		wg.Wait()
		if registerStatus {
			err = s.userRepo.Create(user)
			if err != nil {
				s.logger.Sugar().Error("INSERT USER", zap.String("Error", err.Error()))
			}
			s.logger.Info("CREATE USER SUCCESS", zap.String("Publickey", userInfo.PublicKey))
		}
	}
	return errors.New(200, "User is registed")
}

func (s *UserService) Login(login request.Login) (bool, *model.User, error) {
	return s.userRepo.CheckLogin(login.Password, login.Phone)
}

func (s *UserService) GetUserInfo(c *gin.Context, phone string) (*response.User, error) {
	contractAddress := common.HexToAddress(s.documentContract)
	documentAbi, err := abi.NewAbi(contractAddress, s.ethclient)
	if err != nil {
		s.logger.Sugar().Error(err)
		return nil, err
	}

	hash, err := documentAbi.GetHashUserInfo(&bind.CallOpts{}, phone)
	if err != nil {
		s.logger.Sugar().Error(err)
		return nil, err
	}

	user, err := s.userRepo.GetUserByPhone(phone)
	if err != nil {
		s.logger.Sugar().Error(err)
		return nil, err
	}

	return &response.User{
		ID:          user.ID,
		Name:        user.Name,
		PublicKey:   user.PublicKey,
		CardID:      user.CardID,
		Phone:       user.Phone,
		Gmail:       user.Gmail,
		DateOfBirth: user.DateOfBirth,
		Hash:        string(hash[:]),
	}, nil
}

func (s *UserService) VerifyUser(user request.UserInfo) (bool, error) {
	contractAddress := common.HexToAddress(s.documentContract)
	documentIntance, err := abi.NewAbi(contractAddress, s.ethclient)
	if err != nil {
		s.logger.Sugar().Error(err)
		return false, err
	}
	address := common.HexToAddress(user.PublicKey)
	isVerify, err := documentIntance.VerifyUser(&bind.CallOpts{}, user.Name, user.CardID, user.DateOfBirth, user.Phone, user.Gmail, address)
	return isVerify, err
}
