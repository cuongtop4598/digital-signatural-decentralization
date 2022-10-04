package service

import (
	"digitalsignature/internal/app/model"
	"digitalsignature/internal/app/repository"
	"digitalsignature/internal/app/response"
	"digitalsignature/internal/pkg/abi"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

func NewDocumentService(
	documentRepo *repository.DocumentRepo,
	userRepo *repository.UserRepo,
	client *ethclient.Client,
	log *zap.Logger,
) *DocumentService {
	contractAddress, ok := os.LookupEnv("CONTRACT_ADDRESS")
	if !ok {
		log.Fatal("Can't load contract address")
	}
	return &DocumentService{
		documentRepo:    documentRepo,
		userRepo:        userRepo,
		contractAddress: common.HexToAddress(contractAddress),
		client:          client,
		log:             log,
	}
}

type DocumentService struct {
	documentRepo    *repository.DocumentRepo
	userRepo        *repository.UserRepo
	contractAddress common.Address
	client          *ethclient.Client
	log             *zap.Logger
}

type DocumentEvent struct {
	NumberOfOwnerDocument *big.Int
	CreateAccountSuccess  string
}

func (d *DocumentService) GetAllIsPublic() (*response.ListDocumentResponse, error) {
	result := response.ListDocumentResponse{}
	docs, err := d.documentRepo.AllIsPublic()
	if err != nil {
		d.log.Sugar().Error(err)
		return nil, err
	}
	for _, doc := range docs {
		result.Documents = append(result.Documents, response.DocumentResponse{
			IndexOnchain:    doc.IndexOnchain,
			Owner:           doc.Owner,
			Name:            doc.Name,
			BlockNumber:     doc.BlockNumber,
			BlockHash:       doc.BlockHash,
			TransactionHash: doc.TransactionHash,
			Signature:       doc.Signature,
			Public:          doc.Public,
			TypeFile:        doc.TypeFile,
		})
	}
	return &result, nil
}

func (d *DocumentService) GetDocumentByPhone(phone string) (*response.ListDocumentResponse, error) {
	result := response.ListDocumentResponse{}
	userPublicKey, err := d.userRepo.GetPublickeyByPhone(phone)
	if err != nil {
		d.log.Sugar().Error(err)
		return nil, err
	}
	docs, err := d.documentRepo.AllByOwner(userPublicKey)
	if err != nil {
		d.log.Sugar().Error(err)
		return nil, err
	}
	for _, doc := range docs {
		result.Documents = append(result.Documents, response.DocumentResponse{
			IndexOnchain:    doc.IndexOnchain,
			Owner:           doc.Owner,
			Name:            doc.Name,
			BlockNumber:     doc.BlockNumber,
			BlockHash:       doc.BlockHash,
			TransactionHash: doc.TransactionHash,
			Signature:       doc.Signature,
			Public:          doc.Public,
			TypeFile:        doc.TypeFile,
		})
	}
	return &result, nil
}

func (d *DocumentService) GetDocumentByPublickey(publickey string) ([]model.Document, error) {
	result := response.ListDocumentResponse{}
	docs, err := d.documentRepo.AllByOwner(publickey)
	if err != nil {
		d.log.Sugar().Error(err)
		return nil, err
	}
	for _, doc := range docs {
		result.Documents = append(result.Documents, response.DocumentResponse{
			IndexOnchain:    doc.IndexOnchain,
			Owner:           doc.Owner,
			Name:            doc.Name,
			BlockNumber:     doc.BlockNumber,
			BlockHash:       doc.BlockHash,
			TransactionHash: doc.TransactionHash,
			Signature:       doc.Signature,
			Public:          doc.Public,
			TypeFile:        doc.TypeFile,
		})
	}
	return docs, nil
}

func (d *DocumentService) VerifyDocument(phone string, digest string, documentIndex *big.Int) (bool, error) {
	documentAbi, err := abi.NewDocument(d.contractAddress, d.client)
	if err != nil {
		d.log.Sugar().Error(err)
		return false, err
	}
	isCorrect, err := documentAbi.VerifyDoc(&bind.CallOpts{}, phone, digest, documentIndex)
	return isCorrect, err
}

func (d *DocumentService) GetSignature(phone string, documentIndex *big.Int) ([]byte, error) {
	documentAbi, err := abi.NewDocument(d.contractAddress, d.client)
	if err != nil {
		d.log.Sugar().Error(err)
		return nil, err
	}
	signature, _, err := documentAbi.GetSignature(&bind.CallOpts{}, phone, documentIndex)
	if err != nil {
		d.log.Sugar().Error(err)
		return nil, err
	}
	return signature, nil
}

func (d *DocumentService) StoreDocument(
	fileName string,
	fileType string,
	publickeys string,
	signature string,
	blockHash string,
	transactionHash string,
	blockNumber string,
	documentId int,
	isPublic bool) error {
	document := model.Document{
		ID:              [16]byte{},
		IndexOnchain:    documentId,
		Owner:           publickeys,
		Name:            fileName,
		BlockNumber:     blockNumber,
		BlockHash:       blockHash,
		TransactionHash: transactionHash,
		Signature:       signature,
		Public:          isPublic,
		TypeFile:        fileType,
		CreateAt:        time.Now(),
		UpdateAt:        time.Time{},
		DeleteAt:        time.Time{},
	}
	err := d.documentRepo.Create(&document)
	if err != nil {
		return err
	}
	return nil
}
