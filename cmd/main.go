package main

import (
	"digitalsignature/internal/app"
	"log"

	"github.com/joho/godotenv"
)

var cfgFile string
var s app.Server

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}
	s.Run("development")
}
