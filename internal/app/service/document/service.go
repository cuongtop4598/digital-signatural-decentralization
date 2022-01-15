package document

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

const (
	// geth dumpgenesis
	chainID = 451998                                            // change chainID equals chainID in your genesis.json
	dataDir = "/home/cuongtop/Desktop/DigitalSignaturalNetwork" // path of --datadir setting in your network
)

type DocumentService interface {
	GetUserIdByPublicKey(c *gin.Context, userAddress common.Address) (id string, err error)
	SaveDocument(userID string, passUnlock string) int64
}

type document struct {
	client   *ethclient.Client
	instance *Document
	account  *accounts.Account
	Data     []byte
	chanID   int64
}

func NewDocumentService(client *ethclient.Client, chainID int64) DocumentService {
	return &document{
		client: client,
		chanID: chainID,
	}
}

func (d *document) SetAccount(account *accounts.Account) *document {
	d.account = account
	return d
}

func (d *document) CreateInstance(userAddress common.Address, client *ethclient.Client) *document {
	doc, err := NewDocument(userAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	d.instance = doc
	return d
}

func (d *document) GetUserIdByPublicKey(c *gin.Context, userAddress common.Address) (id string, err error) {
	return d.instance.GetUserID(nil, userAddress)
}

func (d *document) SaveDocument(userID string, passUnlock string) int64 {
	signatura := []byte{}
	// tạm thời gọi với auth là admin
	ks, err := GetKeyStoreAdmin(passUnlock)
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
func (d *document) StoreUser(userInfo *UserInformation, passUnlock string) bool {
	address := common.HexToAddress(userInfo.PublicKey)
	// tạm thời gọi với auth là admin
	ks, err := GetKeyStoreAdmin(passUnlock)
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
