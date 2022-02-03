package config

import (
	"digitalsignature/internal/app/controllers"
	"digitalsignature/internal/app/service"
	"digitalsignature/internal/app/service/document"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	chainID = 451998
)

// NewServices creates a new router services
func NewServices(client *ethclient.Client, DB *gorm.DB, Log *zap.Logger, R *gin.Engine, config *Configuration) *Services {
	return &Services{client, DB, Log, R, config}
}

// Services lets us bind specific services when setting up routes
type Services struct {
	EthClient *ethclient.Client
	DB        *gorm.DB
	Log       *zap.Logger
	R         *gin.Engine
	Config    *Configuration
}

// SetupRoutes instances various repos and services and sets up the routers
func (s *Services) SetupRoutes() {
	controllers.HomeRouter(s.R)
	// Create services handle

	rg := s.R.Group("/v1")

	accountService := service.NewAccountService(s.EthClient, nil, s.Config.Ethereum.Wallets)
	documentService := document.NewDocumentService(s.EthClient, chainID)

	controllers.AccountRouter(accountService, rg)
	controllers.DocumentRouter(documentService, rg)
}
