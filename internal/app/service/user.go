package service

import (
	"context"
	"crypto/sha256"
	"digitalsignature/internal/app/errors"
	"digitalsignature/internal/app/model"
	"digitalsignature/internal/app/repository"
	"digitalsignature/internal/app/request"
	"digitalsignature/internal/app/response"
	"digitalsignature/internal/pkg/abi"
	"digitalsignature/internal/pkg/ethereum"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserService struct {
	ethclient        *ethclient.Client
	userRepo         *repository.UserRepo
	logger           *zap.Logger
	accountSrv       *AdminAccountService
	documentContract string
}

func NewUserService(client *ethclient.Client, userRepo *repository.UserRepo, accountSrv *AdminAccountService, log *zap.Logger) *UserService {
	contractAddress, ok := os.LookupEnv("CONTRACT_ADDRESS")
	if !ok {
		log.Fatal("Can't load document contract")
	}
	return &UserService{
		ethclient:        client,
		userRepo:         userRepo,
		logger:           log,
		accountSrv:       accountSrv,
		documentContract: contractAddress,
	}
}

func (s *UserService) Register(c *gin.Context, userInfo request.UserInfo) error {
	s.ethclient = ethereum.NewClient()
	registerStatus := false
	s.logger.Info("USER_REGISTER", zap.Any("Data", userInfo))
	// TODO: Validate user info including gmail, phone, id card
	// Verify gmail
	{
		// OTP := rand.Int31n(100000)
		// auth := smtp.NewGmailAdminAuth("", "")
		// mail := smtp.NewMail(
		// 	"Digital Signature | User Register",
		// 	[]string{userInfo.Gmail},
		// 	fmt.Sprintf("Your OTP is: %d", OTP),
		// )
		// mail.SendMail(auth)
	}
	// TODO: Check user is exist in database?
	if !s.userRepo.IsExist(userInfo.Phone, userInfo.PublicKey) {
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
		contractAddress := common.HexToAddress(s.documentContract)
		// store hash user info to blockchain
		documentIntance, err := abi.NewDocument(contractAddress, s.ethclient)
		if err != nil {
			s.logger.Sugar().Error(err)
			return err
		}
		address := common.HexToAddress(user.PublicKey)
		txOption := s.accountSrv.GetBindTransactionOptions(s.ethclient)
		s.logger.Sugar().Info("Init store user transaction")
		retry := 0
	RETRY:
		tx, err := documentIntance.StoreUser(txOption, user.Name, user.CardID, user.DateOfBirth, user.Phone, user.Gmail, address)
		if err != nil && retry < 30 {
			s.logger.Sugar().Info("Retry store user onchain!")
			time.Sleep(5 * time.Second)
			retry++
			goto RETRY
		}
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			for {
				receipt, err := s.ethclient.TransactionReceipt(context.Background(), tx.Hash())
				if err != nil {
					s.logger.Info(err.Error())
					s.logger.Info("TRANSACTION_STORE_USER", zap.Any(tx.Hash().String(), "Is Pending"))
					time.Sleep(5 * time.Second)
				} else {
					spew.Dump(receipt)
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
					wg.Done()
					break
				}
			}
		}()
		wg.Wait()
		if registerStatus {
			h := sha256.New()
			user.Password = fmt.Sprintf("%x", h.Sum([]byte(user.Password)))
			err = s.userRepo.Create(user)
			if err != nil {
				s.logger.Sugar().Error("INSERT USER", zap.String("Error", err.Error()))
			}
			s.logger.Info("CREATE USER SUCCESS", zap.String("Publickey", userInfo.PublicKey))
		}
	} else {
		s.logger.Sugar().Error("User is existed")
		return errors.UserIsExisted
	}
	return nil
}

func (s *UserService) Login(login request.Login) (bool, *model.User, error) {
	return s.userRepo.CheckLogin(login.Password, login.Phone)
}

func (s *UserService) GetUserInfo(c *gin.Context, phone string) (*response.User, error) {
	contractAddress := common.HexToAddress(s.documentContract)
	documentAbi, err := abi.NewDocument(contractAddress, s.ethclient)
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
	documentIntance, err := abi.NewDocument(contractAddress, s.ethclient)
	if err != nil {
		s.logger.Sugar().Error(err)
		return false, err
	}
	address := common.HexToAddress(user.PublicKey)
	isVerify, err := documentIntance.VerifyUser(&bind.CallOpts{}, user.Name, user.CardID, user.DateOfBirth, user.Phone, user.Gmail, address)
	return isVerify, err
}
