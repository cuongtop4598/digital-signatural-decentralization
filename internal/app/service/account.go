package service

import (
	"context"
	"errors"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

// This service use to interface with admin account ethereum
type AccountService struct {
	KeyPath  string
	Password string
}

func NewAccountService(keyPath string, password string) *AccountService {
	return &AccountService{
		KeyPath:  keyPath,
		Password: password,
	}
}

func (s *AccountService) GetAminAccount() (*accounts.Account, *keystore.KeyStore, error) {
	os.Chdir(".")
	ks := keystore.NewKeyStore(s.KeyPath, keystore.StandardScryptN, keystore.StandardScryptP)
	accounts := ks.Accounts()
	if len(accounts) > 0 {
		err := ks.Unlock(accounts[0], s.Password)
		if err != nil {
			return nil, nil, err
		}
		return &accounts[0], ks, nil
	}
	return nil, nil, errors.New("No key file in keystore path")
}

func (s *AccountService) BindTransactionOption(account accounts.Account, password string, ks *keystore.KeyStore, client *ethclient.Client) *bind.TransactOpts {
	ks.Unlock(account, password)
	fromAddress := account.Address
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyStoreTransactorWithChainID(ks, account, big.NewInt(451998))
	if err != nil {
		log.Fatal(err)
	}
	auth.From = fromAddress
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = big.NewInt(10000000)
	auth.GasLimit = uint64(0x47b760)

	return auth
}
