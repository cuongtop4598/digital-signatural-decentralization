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
	ar.GET("/download/:filename", dc.Download)
	ar.GET("/list/:publickey", dc.GetDocs)
	ar.GET("/signature", dc.GetSign)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	file, h, err := c.Request.FormFile("filedata")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	publickey := c.Request.FormValue("publickey")
	signature := c.Request.FormValue("signature")

	tmpFile, err := os.Create("static/" + h.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer tmpFile.Close()

	doc := model.Document{
		ID:        uuid.New(),
		Owner:     []string{publickey},
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
	files := utils.SearchFileInPath("static/")
	for _, file := range files {
		if file == ("static/" + name) {
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

	documentId, err := strconv.Atoi(c.Request.FormValue("document_id"))
	if err != nil {
		log.Println("Document Id must be integer value: ", err)
	}

	os.Chdir(".")
	tmpFile, err := os.Create("./static/" + h.Filename + ".pdf")
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
		h.Filename,
		"pdf", []string{publickey},
		signature,
		blockHash,
		txHash,
		blockNumber,
		documentId)
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
			IndexOnchain:    documentId,
			Owner:           []string{publickey},
			Name:            h.Filename,
			BlockNumber:     blockNumber,
			BlockHash:       blockHash,
			TransactionHash: txHash,
			Signature:       signature,
			Public:          false,
			TypeFile:        "pdf",
		},
	})
}

func (dc *DocumentController) GetDocs(c *gin.Context) {
	publickey := c.Param("publickey")
	docs, err := dc.documentService.GetDocumentByPublickey([]string{publickey})
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
