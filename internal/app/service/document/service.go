package document

import (
	"digitalsignature/internal/app/utils"
	"log"

	"github.com/ethereum/go-ethereum/accounts"
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
	account  *accounts.Account
	data     []byte
	chanID   int64
}

func NewDocumentService(data []byte, client *ethclient.Client, userAddress common.Address, account *accounts.Account, chainID int64) DocumentService {
	doc, err := NewDocument(userAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	return &document{
		client:   client,
		instance: doc,
		account:  account,
		data:     data,
		chanID:   chainID,
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
	// tạm thời gọi với auth là admin
	ks, err := utils.GetKeyStoreAdmin()
	if err != nil {
		log.Fatal(err)
	}
	auth, err := CreateAuthForSigning(*d.account, d.client, ks, d.chanID)
	if err != nil {
		log.Fatal(err)
	}
	tnx, err := d.instance.SaveDoc(auth, userID, signatura)

	if err != nil {
		log.Fatal(err)
	}
	_ = tnx // get tnx info here
	return 0
}

// Verify document by using userID, digest, DocID
// userID using for get
func (d *document) VerifyDoc(userID string, digest []byte, DocID string) bool {
	return true
}

// Store user info
func (d *document) StoreUser(userInfo *UserInformation) bool {
	address := common.HexToAddress(userInfo.PublicKey)
	// tạm thời gọi với auth là admin
	ks, err := utils.GetKeyStoreAdmin()
	if err != nil {
		log.Fatal(err)
	}
	auth, err := CreateAuthForSigning(*d.account, d.client, ks, d.chanID)
	if err != nil {
		log.Fatal(err)
	}
	tnx, err := d.instance.StoreUser(auth, userInfo.ID, userInfo.Name, userInfo.IdentityCard, userInfo.DateOfBirth, userInfo.Phone, userInfo.Gmail, address)
	if err != nil {
		log.Fatal(err)
	}
	_ = tnx
	return true
}
