package main

import (
	"digitalsignature/internal/pkg/database"
	"digitalsignature/routers"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	log, _ := zap.NewDevelopment()
	err := godotenv.Load()
	if err != nil {
		log.Error(err.Error())
	}
	schema := "public"
	db := database.NewDBConnection(log, &schema)
	cmd := routers.NewCMD(db, log)
	cmd.MigrateDB()
}
