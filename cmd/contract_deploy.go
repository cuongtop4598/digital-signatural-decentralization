package main

import (
	"context"
	"digitalsignature/config"
	"digitalsignature/internal/app/service/document"
	"digitalsignature/internal/pkg/ethereum"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

const (
	// geth dumpgenesis
	chainID = 451998                                            // change chainID equals chainID in your genesis.json
	dataDir = "/home/cuongtop/Desktop/DigitalSignaturalNetwork" // path of --datadir setting in your network
)

type Keys struct {
	Private string `json:"private"`
}

func GetAccountAdmin(password string) accounts.Account {
	ks := keystore.NewKeyStore(dataDir+"/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	accounts := ks.Accounts()
	fmt.Println("List accounts: ", accounts)
	if len(accounts) > 0 {
		log.Println(accounts[0].Address)
		return accounts[0]
	} else {
		account, _ := ks.NewAccount(password)
		err := os.WriteFile("/wallets/address", []byte(account.Address.Hex()), 0644)
		if err != nil {
			log.Fatal(err)
		}
		privateKey, _ := ks.Export(account, password, password)
		err = os.WriteFile("/wallets/privatekey", privateKey, 0644)
		if err != nil {
			log.Fatal(err)
		}
		return account
	}
}

func main() {
	config, err := config.NewConfig("/config/", "development")
	if err != nil {
		log.Fatal(err)
	}
	client, err := ethereum.NewClient(config.Ethereum)
	if err != nil {
		log.Fatal(err)
	}
	account := GetAccountAdmin(config.Ethereum.Password)
	ks := keystore.NewKeyStore(dataDir+"/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	err = ks.Unlock(account, config.Ethereum.Password)
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := account.Address
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyStoreTransactorWithChainID(ks, account, big.NewInt(chainID))
	if err != nil {
		log.Fatal(err)
	}
	auth.From = account.Address
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(8000000) // in units gas limit: 134217728
	auth.GasPrice = gasPrice
	address, tx, instance, err := document.DeployDocument(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())
	fmt.Println("Deploy Contract Done!")
	_ = instance
}
