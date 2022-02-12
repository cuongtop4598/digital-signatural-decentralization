package document

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AccountSrv interface {
	GetAminAccount() (*accounts.Account, *keystore.KeyStore, error)
}

type DocumentService interface {
}

type document struct {
	accountSrv AccountSrv
	instance   *Document
	data       []byte
	Network    struct {
		ChainID int64
		Client  *ethclient.Client
	}
}

func NewDocumentService(data []byte, client *ethclient.Client, chainID int64, accountSrv AccountSrv, userAddress common.Address, account *accounts.Account) DocumentService {
	doc, err := NewDocument(userAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	return &document{
		accountSrv: accountSrv,
		instance:   doc,
		data:       data,
		Network: struct {
			ChainID int64
			Client  *ethclient.Client
		}{ChainID: chainID, Client: client},
	}
}

// Verify document by using userID, digest, DocID
// userID using for get
func (d *document) VerifyDoc(userID string, digest []byte, DocID string) bool {
	return true
}

// Store user info
func (d *document) StoreUser(userInfo *UserInformation) bool {

	adminAccount, ks, err := d.accountSrv.GetAminAccount()
	if err != nil {
		log.Fatal(err)
	}
	auth, err := createAuthForSigning(*adminAccount, d.Network.Client, ks, d.Network.ChainID)
	if err != nil {
		log.Fatal(err)
	}
	tnx, err := d.instance.StoreUser(auth, userInfo.ID, userInfo.Name, userInfo.IdentityCard, userInfo.DateOfBirth, userInfo.Phone, userInfo.Gmail, userInfo.PublicKey)
	if err != nil {
		log.Fatal(err)
	}
	_ = tnx
	return true
}

func createAuthForSigning(account accounts.Account, client *ethclient.Client, ks *keystore.KeyStore, chainID int64) (auth *bind.TransactOpts, err error) {

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
