package document

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts"
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
	instance *Document
}

func NewDocumentService(client *ethclient.Client, userAddress common.Address) DocumentService {
	doc, err := NewDocument(userAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	return &document{
		instance: doc,
	}
}

// Verify document by using userID, digest, DocID
// userID using for get
func (d *document) VerifyDoc(userID string, digest []byte, DocID string) bool {
	return true
}
