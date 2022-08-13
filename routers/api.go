package routers

import (
	"digitalsignature/internal/app/controllers"
	"digitalsignature/internal/app/repository"
	"digitalsignature/internal/app/service"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// NewRouter creates a new router services
func NewRouter(client *ethclient.Client, DB *gorm.DB, RD *redis.Client, Log *zap.Logger, R *gin.Engine) *Router {
	return &Router{client, DB, RD, Log, R}
}

type Router struct {
	EthClient *ethclient.Client
	DB        *gorm.DB
	Redis     *redis.Client
	Log       *zap.Logger
	R         *gin.Engine
}

// SetupRoutes instances various repos and services and sets up the routers
func (r *Router) SetupRoutes() {
	userRepo := repository.NewUserRepository(r.DB, r.Log)
	docRepo := repository.NewDocumentRepo(r.DB, r.Log)

	adminAccountSrv := service.NewAdminAccountService()
	documentSrv := service.NewDocumentService(docRepo, userRepo, r.EthClient, r.Log)
	userSrv := service.NewUserService(r.EthClient, userRepo, adminAccountSrv, r.Log)

	rg := r.R.Group("/v1")
	controllers.HomeRouter(r.R)
	controllers.UserRouter(userSrv, rg)
	controllers.DocumentRouter(documentSrv, docRepo, userRepo, rg)
}
