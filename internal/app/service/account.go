package service

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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
	keyJson, err := ks.Export(account, password, password)
	if err != nil {
		log.Fatal(err)
	}
	key, err := keystore.DecryptKey(keyJson, password)
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
	return &bind.TransactOpts{
		From:     fromAddress,
		Nonce:    big.NewInt(int64(nonce)),
		Signer:   keySigner(big.NewInt(451998), key.PrivateKey),
		GasPrice: big.NewInt(200),
		GasLimit: uint64(100000000),
		Context:  context.TODO(),
		NoSend:   true,
	}

}

func keySigner(chainID *big.Int, key *ecdsa.PrivateKey) (signerfn bind.SignerFn) {
	signerfn = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		keyAddr := crypto.PubkeyToAddress(key.PublicKey)
		if address != keyAddr {
			return nil, errors.New("not authorized to sign this account")
		}
		return types.SignTx(tx, types.NewEIP155Signer(chainID), key)
	}
	return
}
