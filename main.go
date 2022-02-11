package main

import (
	"digitalsignature/internal/app"
	"log"
)

var server app.Server

func main() {
	err := server.Run("development")
	log.Fatal("Running server error: ", err)
}
