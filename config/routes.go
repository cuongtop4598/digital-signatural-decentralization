package config

import (
	"digitalsignature/internal/app/controllers"
	"digitalsignature/internal/app/service"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// NewServices creates a new router services
func NewServices(client *ethclient.Client, DB *pg.DB, RD *redis.Client, Log *zap.Logger, R *gin.Engine, config *Configuration) *Services {
	return &Services{client, DB, RD, Log, R, config}
}

// Services lets us bind specific services when setting up routes
type Services struct {
	EthClient *ethclient.Client
	DB        *pg.DB
	Redis     *redis.Client
	Log       *zap.Logger
	R         *gin.Engine
	Config    *Configuration
}

// SetupRoutes instances various repos and services and sets up the routers
func (s *Services) SetupRoutes() {

	// Create services handle
	controllers.HomeRouter(s.R)
	rg := s.R.Group("/v1")

	accountService := service.NewAccountService(s.EthClient, nil, s.Config.Ethereum.Wallets)
	controllers.AccountRouter(accountService, rg)
}
