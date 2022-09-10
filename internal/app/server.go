package app

import (
	"context"
	"digitalsignature/internal/app/middleware"
	"digitalsignature/internal/pkg/database"
	"digitalsignature/internal/pkg/ethereum"
	"os"
	"time"

	"digitalsignature/routers"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// Server holds all the routes and their services
type Server struct{}

// Run runs our API server
func (server *Server) Run(env string) error {
	gin.ForceConsoleColor()
	r := gin.Default()
	headerPolicies := cors.DefaultConfig()
	headerPolicies.AllowOrigins = []string{
		"http://localhost:3000",
		"http://localhost:3001",
		"http://localhost:3002",
		"http://192.168.78.2:3000",
		"http://192.168.78.2:3001",
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
	r.Use(middleware.CORSMiddleware())
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
	schema := "public"
	db := database.NewDBConnection(log, &schema)

	client := ethereum.NewClient()

	currentBlockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Sugar().Info("Block number: ", currentBlockNumber)
	defer tracer.Stop()
	defer log.Sync()

	rsDefault := &routers.Router{
		EthClient: client,
		DB:        db,
		Log:       log,
		R:         r,
	}
	// setup router
	rsDefault.SetupRoutes()

	port, ok := os.LookupEnv("BACKEND_PORT")
	if !ok {
		port = "8080"
	}
	return r.Run(":" + port)
}
