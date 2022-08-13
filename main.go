package main

import (
	"digitalsignature/internal/app"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var server app.Server

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}
	env := os.Getenv("GIN_MODE")
	if env == "" {
		env = "release"
	}
	server.Run(env)
}
