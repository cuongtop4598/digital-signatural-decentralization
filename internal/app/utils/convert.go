package utils

import (
	"digitalsignature/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

const (
	// geth dumpgenesis
	chainID = 451998                                            // change chainID equals chainID in your genesis.json
	dataDir = "/home/cuongtop/Desktop/DigitalSignaturalNetwork" // path of --datadir setting in your network
)

func GetKeyStoreAdmin() (*keystore.KeyStore, error) {
	config, err := config.NewConfig("/config/", "development")
	if err != nil {
		return nil, err
	}
	account := GetAccountAdmin(config.Ethereum.Password)
	ks := keystore.NewKeyStore(dataDir+"/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	err = ks.Unlock(account, config.Ethereum.Password)
	if err != nil {
		return nil, err
	}
	return ks, nil
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

func Json2struct(jsonfile string, dst interface{}) error {
	file, err := ioutil.ReadFile(jsonfile)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(file), &dst)
	if err != nil {
		return err
	}
	return nil
}
