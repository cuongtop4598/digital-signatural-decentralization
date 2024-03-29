package controllers

import (
	"digitalsignature/internal/app/middleware"
	"digitalsignature/internal/app/model"
	"digitalsignature/internal/app/repository"
	"digitalsignature/internal/app/request"
	"digitalsignature/internal/app/response"
	"digitalsignature/internal/app/service"
	"digitalsignature/internal/app/utils"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DocumentController struct {
	documentService *service.DocumentService
	documentRepo    *repository.DocumentRepo
	userRepo        *repository.UserRepo
}

func DocumentRouter(docService *service.DocumentService, documentRepo *repository.DocumentRepo, userRepo *repository.UserRepo, r *gin.RouterGroup) {
	dc := DocumentController{
		documentService: docService,
		documentRepo:    documentRepo,
		userRepo:        userRepo,
	}
	ar := r.Group("/document", middleware.AuthMiddleware())
	ar.POST("/savesign", dc.SaveSignedDocument)
	ar.POST("/upload", dc.Upload)
	ar.GET("/list", dc.GetAllIsPublic)
	ar.GET("/list-multi/:publickey", dc.GetAllMultiSignerDocument)
	ar.GET("/list/:publickey", dc.GetDocs)
	ar.GET("/signature", dc.GetSign)
	ar.GET("/check/:digest", dc.CheckIsExisted)
	ar.POST("/save-multiple", dc.SaveMultiSignedDocument)
	ar.POST("/update-multiple", dc.UpdateMultiSignedDocument)

	publicRouter := r.Group("/document")
	publicRouter.GET("/download/filename/:filename/owner/:owner", dc.Download)
}

func (dc *DocumentController) GetSign(c *gin.Context) {
	getSignRequest := request.GetSignRequest{}
	err := c.BindJSON(&getSignRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sign, err := dc.documentService.GetSignature(getSignRequest.Phone, big.NewInt(int64(getSignRequest.Number)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(len(sign))
	c.JSON(http.StatusOK, gin.H{"length": len(sign), "sign": sign})
}

func (dc *DocumentController) Upload(c *gin.Context) {

	err := c.Request.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		log.Println(err)
	}
	file, h, err := c.Request.FormFile("filedata")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	publickey := c.Request.FormValue("publickey")
	signature := c.Request.FormValue("signature")

	path := fmt.Sprintf("static/" + publickey)
	log.Println(path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		log.Println(err)
	}
	tmpFile, err := os.Create(path + h.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer tmpFile.Close()

	doc := model.Document{
		ID:        uuid.New(),
		Owner:     publickey,
		Name:      h.Filename,
		TypeFile:  "pdf",
		Signature: signature,
		Public:    true,
		CreateAt:  time.Now(),
	}
	err = dc.documentRepo.Create(&doc)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = io.Copy(tmpFile, file)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "update load file success"})
}

func (dc *DocumentController) Download(c *gin.Context) {
	name := c.Param("filename")
	owner := c.Param("owner")
	files := utils.SearchFileInPath("./static/" + owner + "/")
	for _, file := range files {
		if file == ("static\\" + owner + "\\" + name) {
			isPublic, err := dc.documentRepo.IsPublic(name)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
			if isPublic {
				c.File(file)
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "file's not public"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "file not found"})
}

type VerifyDocRequest struct {
	Phone  string `json:"phone"`
	Digest string `json:"digest"`
	DocNum int64  `json:"doc_number"`
}

func (dc *DocumentController) SaveSignedDocument(c *gin.Context) {
	err := c.Request.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		log.Println("Parse multi part form", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	file, h, err := c.Request.FormFile("filedata")
	if err != nil {
		log.Println("Get file error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	publickey := c.Request.FormValue("publickey")
	signature := c.Request.FormValue("signature")
	blockNumber := c.Request.FormValue("block_number")
	blockHash := c.Request.FormValue("block_hash")
	txHash := c.Request.FormValue("transaction_hash")
	isPublic, err := strconv.ParseBool(c.Request.FormValue("public"))
	if err != nil {
		log.Println(err)
	}
	documentId, err := strconv.ParseInt(c.Request.FormValue("document_id"), 10, 64)
	if err != nil {
		log.Println("Document Id must be integer value: ", err)
	}

	os.Chdir(".")
	path := fmt.Sprintf("./static/" + publickey)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, os.ModePerm)
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + "-" + h.Filename
	tmpFile, err := os.Create(path + "/" + fileName)
	if err != nil {
		log.Println("Save file error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	defer tmpFile.Close()
	_, err = io.Copy(tmpFile, file)
	if err != nil {
		log.Println("copy file error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	err = dc.documentService.StoreDocument(
		fileName,
		"pdf",
		publickey,
		signature,
		blockHash,
		txHash,
		blockNumber,
		int(documentId),
		isPublic)
	if err != nil {
		log.Println("create document error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "OK",
		"data": response.DocumentResponse{
			IndexOnchain:    int(documentId),
			Owner:           publickey,
			Name:            h.Filename,
			BlockNumber:     blockNumber,
			BlockHash:       blockHash,
			TransactionHash: txHash,
			Signature:       signature,
			Public:          isPublic,
			TypeFile:        "pdf",
		},
	})
}

func (dc *DocumentController) GetDocs(c *gin.Context) {
	publickey := c.Param("publickey")
	docs, err := dc.documentService.GetDocumentByPublickey(publickey)
	if err != nil {
		log.Println("get list document error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "OK",
		"data":    docs,
	})
}

func (dc *DocumentController) GetAllIsPublic(c *gin.Context) {
	docs, err := dc.documentService.GetAllIsPublic()
	if err != nil {
		log.Println("get list document error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "OK",
		"data":    docs,
	})
}

func (dc *DocumentController) SaveMultiSignedDocument(c *gin.Context) {
	err := c.Request.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		log.Println("Parse multi part form", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	file, h, err := c.Request.FormFile("filedata")
	if err != nil {
		log.Println("Get file error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	creater := c.Request.FormValue("publickey")
	signature := c.Request.FormValue("signature")
	blockNumber := c.Request.FormValue("block_number")
	blockHash := c.Request.FormValue("block_hash")
	txHash := c.Request.FormValue("transaction_hash")
	digest := c.Request.FormValue("digest")
	partnerB := c.Request.FormValue("partner_b")
	documentTimestamp := c.Request.FormValue("document_timestamp")

	isPublic, err := strconv.ParseBool(c.Request.FormValue("public"))
	if err != nil {
		log.Println(err)
	}

	documentId, err := strconv.ParseInt(c.Request.FormValue("document_index"), 10, 64)
	if err != nil {
		log.Println("Document Id must be integer value: ", err)
	}

	os.Chdir(".")
	path := fmt.Sprintf("./static/" + creater)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, os.ModePerm)
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + "-" + h.Filename
	tmpFile, err := os.Create(path + "/" + fileName)
	if err != nil {
		log.Println("Save file error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	defer tmpFile.Close()
	_, err = io.Copy(tmpFile, file)
	if err != nil {
		log.Println("copy file error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	err = dc.documentService.StoreMultiSignDocument(
		fileName,
		"pdf",
		creater,
		partnerB,
		blockHash,
		digest,
		signature,
		txHash,
		blockNumber,
		int(documentId),
		documentTimestamp,
		isPublic)
	if err != nil {
		log.Println("create document error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "OK",
		"data": response.MultiSignerDocumentResponse{
			IndexOnchain:    int(documentId),
			Creater:         creater,
			PartnerA:        creater,
			PartnerB:        partnerB,
			Name:            h.Filename,
			BlockNumber:     blockNumber,
			BlockHash:       blockHash,
			TransactionHash: txHash,
			SignatureA:      signature,
			Digest:          digest,
			SignatureB:      documentTimestamp,
			Public:          isPublic,
			TypeFile:        "pdf",
		},
	})
}

func (dc *DocumentController) UpdateMultiSignedDocument(c *gin.Context) {
	err := c.Request.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		log.Println("Parse multi part form", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	file, h, err := c.Request.FormFile("filedata")
	if err != nil {
		log.Println("Get file error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	signature := c.Request.FormValue("signature_b")
	signDate := c.Request.FormValue("sign_date")
	complete, _ := strconv.ParseBool(c.Request.FormValue("success"))
	partnerB := c.Request.FormValue("partner_b")
	id := c.Request.FormValue("id")
	os.Chdir(".")
	path := fmt.Sprintf("./static/" + partnerB)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, os.ModePerm)
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + "-" + h.Filename
	tmpFile, err := os.Create(path + "/" + fileName)
	if err != nil {
		log.Println("Save file error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	defer tmpFile.Close()
	_, err = io.Copy(tmpFile, file)
	if err != nil {
		log.Println("copy file error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	err = dc.documentService.UpdatePartnerInMultiSign(id, complete, signature, signDate)
	if err != nil {
		log.Println("create document error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "OK",
	})
}

func (dc *DocumentController) GetAllMultiSignerDocument(c *gin.Context) {
	publickey := c.Param("publickey")
	docs, err := dc.documentService.GetMultiDocumentByPartner(publickey)
	if err != nil {
		log.Println("get list document error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "OK",
		"data":    docs,
	})
}

func (dc *DocumentController) CheckIsExisted(c *gin.Context) {
	digest := c.Param("digest")
	document := dc.documentService.GetMultiByDigest(digest)
	if document != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":       http.StatusOK,
			"is_existed": true,
			"document":   document,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":       http.StatusOK,
			"is_existed": false,
			"document":   "",
		})
	}
}
