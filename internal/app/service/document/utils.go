package document

import (
	"context"
	"log"
	"math/big"

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
