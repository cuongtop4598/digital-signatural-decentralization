package service

import (
	"digitalsignature/internal/app/repository/offchain"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type AccountService interface {
	CreateNewWallet()
}
type Account struct {
	Client      *ethclient.Client
	AccountRepo *offchain.AccountRepo
	walletsDir  string
}

func (a *Account) CreateNewAccount(c *gin.Context) ([]byte, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	return privateKeyBytes, nil
}

func (a *Account) CreateNewKeyStores(password string, c *gin.Context) (publickey string, err error) {
	ks := keystore.NewKeyStore(a.walletsDir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		return "", err
	}
	return account.Address.Hex(), nil
}

func (a *Account) ImportKeyStores(password string, keyPath string, c gin.Context) {
	file := keyPath
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}
