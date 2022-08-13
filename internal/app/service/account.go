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

func NewAdminAccountService() *AdminAccountService {
	keyPath, _ := os.LookupEnv("KEYSTORE_PATH")
	password, _ := os.LookupEnv("PASSWORD_UNLOCK")
	privatekey, ok := os.LookupEnv("PRIVATE_KEY")
	if !ok {
		privatekey = "0x07535d6917bdf7ca39da2f477b7f96d350672a2549f546572d905fb553fc9988"
	}
	private, err := ethereumCrypto.HexToECDSA(privatekey)
	if err != nil {
		log.Fatal(err)
	}
	pubkey, ok := private.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	return &AdminAccountService{
		KeyPath:    keyPath,
		Password:   password,
		Publickey:  pubkey,
		PrivateKey: private,
	}
}

func (s *AdminAccountService) GetBindTransactionOptions(client *ethclient.Client) *bind.TransactOpts {
	fromAddress := ethereumCrypto.PubkeyToAddress(*s.Publickey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		s.log.Sugar().Error(err)
	}
	gasLimit := uint64(4712388) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		s.log.Sugar().Error(err)
	}
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
		GasPrice: gasPrice,
		GasLimit: gasLimit,
		Context:  context.Background(),
	}
}

func (s *AdminAccountService) SignTransaction(tx *types.Transaction, client *ethclient.Client) (*types.Transaction, error) {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), s.PrivateKey)
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
	auth.GasPrice = big.NewInt(1000000000)
	auth.GasLimit = uint64(4712388)
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
