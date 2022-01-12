package document

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type DocumentService interface {
	GetUserIdByPublicKey(c *gin.Context, userAddress common.Address) (id string, err error)
}

type document struct {
	client   *ethclient.Client
	instance *Document
	data     []byte
}

func NewDocumentService(data []byte, client *ethclient.Client, userAddress common.Address) DocumentService {
	doc, err := NewDocument(userAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	return &document{
		client:   client,
		instance: doc,
		data:     data,
	}
}

func (d *document) GetUserIdByPublicKey(c *gin.Context, userAddress common.Address) (id string, err error) {
	return d.instance.GetUserID(nil, userAddress)
}

// func (d *document) GetPublicKeyByUserId(c *gin.Context, userID string) (publicKey string, err error) {
// 	return d.instance.GetPublicKey()
// }

func (d *document) SaveDocument(userID string) int64 {
	signatura := []byte{}
	tnx, err := d.instance.SaveDoc(nil, userID, signatura)
	if err != nil {
		log.Fatal(err)
	}
	_ = tnx // get tnx info here
	return 0
}

// more function here
