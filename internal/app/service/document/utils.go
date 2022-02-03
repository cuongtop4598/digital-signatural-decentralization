package document

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CreateAuthForSigning(account accounts.Account, client *ethclient.Client, ks *keystore.KeyStore, chainID int64) (auth *bind.TransactOpts, err error) {

	fromAddress := account.Address
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err = bind.NewKeyStoreTransactorWithChainID(ks, account, big.NewInt(chainID))
	if err != nil {
		log.Fatal(err)
	}
	auth.From = account.Address
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(8000000) // in units gas limit: 134217728
	auth.GasPrice = gasPrice
	return auth, nil
}

func GetKeyStoreAdmin(password string) (*keystore.KeyStore, error) {
	account := GetAccountAdmin(password)
	ks := keystore.NewKeyStore(dataDir+"/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	err := ks.Unlock(account, password)
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
