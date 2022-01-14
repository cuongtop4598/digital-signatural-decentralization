package app

import (
	"digitalsignature/config"
	"digitalsignature/internal/app/migration"
	"digitalsignature/internal/pkg/database"
	"digitalsignature/internal/pkg/ethereum"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// configuration hold all config of server
var configuration *config.Configuration

// Server holds all the routes and their services
type Server struct{}

// Run runs our API server
func (server *Server) Run(env string) error {
	configuration = config.Load(env)
	fmt.Println(configuration.Ethereum.Wallets)
	r := gin.Default()

	headerPolicies := cors.DefaultConfig()
	headerPolicies.AllowOrigins = []string{
		"http://localhost:3000",
	}
	headerPolicies.AllowHeaders = []string{
		"Access-Control-Allow-Credentials",
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Authorization",
		"Access-Control-Allow-Origin",
	}
	headerPolicies.AllowCredentials = true
	headerPolicies.MaxAge = (24 * time.Hour)

	r.Use(cors.New(headerPolicies))
	r.Use(gintrace.Middleware(""))

	log, _ := zap.NewDevelopment()
	if env != "debug" {
		log, _ = zap.NewProduction()
	} else {
		r.Use(ginzap.Ginzap(log, time.RFC3339, true))
		r.Use(ginzap.RecoveryWithZap(log, true))
		tracer.Start(
			tracer.WithEnv(env),
			tracer.WithService(""),
		)
	}

	db, err := database.GetConnection()
	if err != nil {
		log.Sugar().Errorf("Connect to database fail ", err)
	}
	err = migration.Migrate(db)
	if err != nil {
		log.Sugar().Errorf("Migrate fail ", err)
	}
	//rd := redis.GetConnection()

	client, err := ethereum.NewClient(configuration.Ethereum)
	if err != nil {
		log.Sugar().Error(err)
	}

	defer tracer.Stop()
	defer log.Sync()

	rsDefault := &config.Services{
		EthClient: client,
		DB:        db,
		Redis:     nil,
		Log:       log,
		R:         r,
		Config:    configuration,
	}
	// setup router
	rsDefault.SetupRoutes()

	host := configuration.Server.Host
	port := configuration.Server.Port
	return r.Run(host + ":" + port)
}
