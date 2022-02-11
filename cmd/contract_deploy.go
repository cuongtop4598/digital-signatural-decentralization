package main

import (
	"context"
	"digitalsignature/internal/app/service"
	"digitalsignature/internal/app/service/document"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	chainID = int64(451998)
)

func main() {
	client, err := ethclient.Dial("http://27.72.105.169:10199")
	if err != nil {
		log.Fatal(err)
	}
	// TODO: deploy contract
	// get admin account
	accountSrv := service.NewAccountService("../wallets/keystore/", "123456")
	account, ks, err := accountSrv.GetAminAccount()
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyStoreTransactorWithChainID(ks, *account, big.NewInt(chainID))
	if err != nil {
		log.Fatal(err)
	}
	fromAddress := account.Address
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	auth.From = account.Address
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(1)        // in wei
	auth.GasLimit = uint64(100000000) // in units gas limit: 134217728
	auth.GasPrice = big.NewInt(100)
	contractAddress, _, _, err := document.DeployDocument(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Contract address", contractAddress)
	fmt.Println("Deploy done!")
}
