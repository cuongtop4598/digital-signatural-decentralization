package service

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
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
