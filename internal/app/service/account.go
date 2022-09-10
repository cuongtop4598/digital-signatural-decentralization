package service

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ethereumCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

// This service use to interface with admin account ethereum
type AdminAccountService struct {
	KeyPath    string
	Password   string
	Publickey  *ecdsa.PublicKey
	PrivateKey *ecdsa.PrivateKey
	log        *zap.Logger
}

func NewAdminAccountService(log *zap.Logger) *AdminAccountService {
	keyPath, _ := os.LookupEnv("KEYSTORE_PATH")
	password, _ := os.LookupEnv("PASSWORD_UNLOCK")
	privatekey, ok := os.LookupEnv("PRIVATE_KEY")
	if !ok {
		log.Fatal("Can't load privatekey")
	}
	private, err := ethereumCrypto.HexToECDSA(privatekey)
	if err != nil {
		log.Sugar().Error(err)
	}
	pubkey, ok := private.Public().(*ecdsa.PublicKey)
	log.Sugar().Info("Publickey: ", ethereumCrypto.PubkeyToAddress(*pubkey))
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	return &AdminAccountService{
		KeyPath:    keyPath,
		Password:   password,
		Publickey:  pubkey,
		PrivateKey: private,
		log:        log,
	}
}

func (s *AdminAccountService) GetBindTransactionOptions(client *ethclient.Client) *bind.TransactOpts {
	fromAddress := ethereumCrypto.PubkeyToAddress(*s.Publickey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		s.log.Sugar().Error(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		s.log.Sugar().Error(err)
	}
	s.log.Sugar().Info("Gas price :", gasPrice)
	return &bind.TransactOpts{
		From:  fromAddress,
		Nonce: big.NewInt(int64(nonce)),
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != fromAddress {
				return nil, errors.New("not authorized to sign this account")
			}
			signedTx, err := s.SignTransaction(tx, client)
			if err != nil {
				return nil, err
			}
			return signedTx, nil
		},
		Value:    big.NewInt(0),
		GasPrice: gasPrice,
		Context:  context.Background(),
	}
}

func (s *AdminAccountService) SignTransaction(tx *types.Transaction, client *ethclient.Client) (*types.Transaction, error) {
	chainID, ok := os.LookupEnv("CHAIN_ID")
	if !ok {
		chainID = "451998"
	}
	chainId, _ := strconv.ParseInt(chainID, 10, 32)
	fmt.Println("Chain id: ", chainId)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(chainId)), s.PrivateKey)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}

func (s *AdminAccountService) BindTransactionOption(account accounts.Account, password string, ks *keystore.KeyStore, client *ethclient.Client) *bind.TransactOpts {
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
	auth.GasPrice = big.NewInt(1000)
	auth.GasLimit = uint64(8000000)
	return auth
}

func (s *AdminAccountService) BindTransactionOptionWithGasLimit(account accounts.Account, password string, ks *keystore.KeyStore, client *ethclient.Client, gas uint64) *bind.TransactOpts {
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
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.From = fromAddress
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = gas
	return auth
}
