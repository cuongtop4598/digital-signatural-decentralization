package app

import (
	"digitalsignature/config"
	"digitalsignature/internal/pkg/ethereum"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// Configuration hold all config of server
var Configuration *config.Configuration

// Server holds all the routes and their services
type Server struct{}

// Run runs our API server
func (server *Server) Run(env string) error {
	Configuration = config.Load(env)

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

	//db := database.GetConnection(log)
	//rd := redis.GetConnection()

	client, err := ethereum.NewClient(Configuration.Ethereum)
	if err != nil {
		log.Sugar().Error(err)
	}

	defer tracer.Stop()
	defer log.Sync()

	rsDefault := &config.Services{
		EthClient: client,
		DB:        nil,
		Redis:     nil,
		Log:       log,
		R:         r,
	}

	rsDefault.SetupRoutes()
	host := Configuration.Server.Host
	port := Configuration.Server.Port
	return r.Run(host + ":" + port)
}
