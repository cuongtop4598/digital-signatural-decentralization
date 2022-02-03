package main

import (
	"digitalsignature/config"
	"digitalsignature/internal/app/service/ipfs"
	"fmt"

	"go.uber.org/zap"
)

// var server app.Server

// func main() {
// 	err := server.Run("development")
// 	log.Fatal("Running server error: ", err)
// }

//=========== TEST ==============
func main() {
	log, _ := zap.NewProduction()
	c, err := config.NewConfig("config", "dev")
	if err != nil {
		log.Fatal(err.Error())
	}
	ipfs := ipfs.NewIpfsService(*c.Ipfs, log)
	cd, err := ipfs.AddFile([]byte("hello cuong nha"))
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(cd)
	data, err := ipfs.GetFile(cd, "hello.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%s", data)
}
