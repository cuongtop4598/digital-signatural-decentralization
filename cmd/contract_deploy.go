// package main

// import (
// 	"context"
// 	"digitalsignature/internal/app/service"
// 	"digitalsignature/internal/app/service/document"
// 	"fmt"
// 	"log"
// 	"math/big"

// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// const (
// 	endpoint = "http://27.72.105.169:10198"
// 	chainID  = int64(451998)
// 	keystore = "../wallets/keystore/"
// 	password = "123456"
// )

// func main() {
// 	client, err := ethclient.Dial(endpoint)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	accountSrv := service.NewAccountService(keystore, password)
// 	account, ks, err := accountSrv.GetAminAccount()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	auth, err := bind.NewKeyStoreTransactorWithChainID(ks, *account, big.NewInt(chainID))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fromAddress := account.Address
// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	auth.From = account.Address
// 	auth.Nonce = big.NewInt(int64(nonce))
// 	auth.Value = big.NewInt(0)
// 	auth.GasLimit = uint64(0x47b760)
// 	auth.GasPrice = big.NewInt(1000)

// 	contractAddress, _, _, err := document.DeployDocument(auth, client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Contract address", contractAddress)
// 	fmt.Println("Deploy done!")
// }
