package service

import (
	"digitalsignature/internal/app/model"
	"digitalsignature/internal/app/repository"
	"digitalsignature/internal/app/request"
	"digitalsignature/internal/app/service/document"
	"log"
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
	BindTransactionOption(account accounts.Account, client *ethclient.Client) *bind.TransactOpts
}
type UserService struct {
	ethclient  *ethclient.Client
	userRepo   *repository.UserRepo
	logger     *zap.Logger
	accountSrv AccountSrv
}

func NewUserService(client *ethclient.Client, userRepo *repository.UserRepo, accountSrv AccountSrv, log *zap.Logger) *UserService {
	return &UserService{
		ethclient:  client,
		userRepo:   userRepo,
		logger:     log,
		accountSrv: accountSrv,
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
	adminAccount, _, err := s.accountSrv.GetAminAccount()
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress("0xF78540428B81A25B7DBe04c19767F47A6f95FEe3")
	// store hash user info to blockchain
	documentIntance, err := document.NewDocument(contractAddress, s.ethclient)
	if err != nil {
		log.Fatal(err)
	}

	userAddress := common.HexToAddress(user.PublicKey)
	tnxOption := s.accountSrv.BindTransactionOption(*adminAccount, s.ethclient)
	txn, err := documentIntance.StoreUser(tnxOption, user.ID.String(), user.Name, user.CardID, user.DateOfBirth, user.Phone, user.Gmail, userAddress)
	if err != nil {
		log.Fatal(err)
	}
	s.logger.Info(txn.Hash().String())
	s.logger.Info("created user", zap.String("publickey", userInfo.PublicKey))
	return nil
}
