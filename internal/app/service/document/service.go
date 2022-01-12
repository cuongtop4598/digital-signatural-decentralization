package document

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type DocumentService interface {
	LoadDocumentContract(c *gin.Context) (instance *Document, err error)
}

type document struct {
	client      *ethclient.Client
	userAddress common.Address
	data        []byte
}

func NewDocumentService(data []byte, client *ethclient.Client, userAddress common.Address) DocumentService {
	return &document{
		client:      client,
		userAddress: userAddress,
		data:        data,
	}
}

func (d *document) LoadDocumentContract(c *gin.Context) (instance *Document, err error) {
	return NewDocument(d.userAddress, d.client)
}
