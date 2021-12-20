package app

import (
	"os"
	"time"

	"depocket.io/config"
	"depocket.io/pkg/database"
	"depocket.io/pkg/redis"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// Server holds all the routes and their services
type Server struct{}

// Run runs our API server
func (server *Server) Run(env string) error {
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

	db := database.GetConnection(log)
	rd := redis.GetConnection()

	eth_endpoint, _ := os.LookupEnv("CHAIN_ENDPOINT")
	client, err := ethclient.Dial(eth_endpoint)

	defer tracer.Stop()

	if err != nil {
		log.Sugar().Error(err)
	}

	defer log.Sync()

	rsDefault := &config.Services{
		EthClient: client,
		DB:        db,
		Redis:     rd,
		Log:       log,
		R:         r,
	}
	rsDefault.SetupRoutes()

	port, ok := os.LookupEnv("DIGITAL_SIGNATURE_BACKEND_PORT")
	if !ok {
		port = "8080"
	}

	return r.Run(":" + port)
}
