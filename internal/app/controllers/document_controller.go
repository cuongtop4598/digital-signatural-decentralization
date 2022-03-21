package controllers

import (
	"digitalsignature/internal/app/model"
	"digitalsignature/internal/app/repository"
	"digitalsignature/internal/app/service/document"
	"digitalsignature/internal/app/utils"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DocumentController struct {
	documentSrv  document.DocumentService
	documentRepo repository.DocumentRepository
	userRepo     repository.UserRepo
}

func DocumentRouter(docService document.DocumentService, documentRepo repository.DocumentRepository, userRepo repository.UserRepo, r *gin.RouterGroup) {
	dc := DocumentController{
		documentSrv:  docService,
		documentRepo: documentRepo,
		userRepo:     userRepo,
	}
	ar := r.Group("/document")
	ar.POST("/savesign", dc.SaveSign)
	ar.POST("/upload", dc.Upload)
	ar.GET("/download/:filename", dc.Download)
	ar.POST("/verify", dc.Verify)
	ar.GET("/list/:publickey", dc.GetDocs)
	ar.GET("/signature", dc.GetSign)
}

type GetSignRequest struct {
	Phone  string `json:"phone"`
	Number int    `json:"number"`
}

func (dc *DocumentController) GetSign(c *gin.Context) {
	getSignRequest := GetSignRequest{}
	err := c.BindJSON(&getSignRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sign, err := dc.documentSrv.GetSignature(getSignRequest.Phone, big.NewInt(int64(getSignRequest.Number)))
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
	file, h, err := c.Request.FormFile("doc")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	publickey := c.Request.FormValue("publickey")
	signature := c.Request.FormValue("signature")

	tmpFile, err := os.Create("static/" + h.Filename + ".pdf")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer tmpFile.Close()

	doc := model.Document{
		DocID:     uuid.New(),
		Owner:     publickey,
		Name:      h.Filename,
		Type:      "pdf",
		Signature: signature,
		Path:      "static/",
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

func (dc *DocumentController) Verify(c *gin.Context) {
	verify := VerifyDocRequest{}
	err := c.BindJSON(&verify)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	isTrue, err := dc.documentSrv.VerifyDocument(verify.Phone, verify.Digest, big.NewInt(verify.DocNum))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	if isTrue {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "OK",
			"data":    "True",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "OK",
			"data":    "False",
		})
		return
	}
}

func (dc *DocumentController) SaveSign(c *gin.Context) {
	err := c.Request.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		log.Println("parse multi part form", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	file, h, err := c.Request.FormFile("doc")
	if err != nil {
		log.Println("get file error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	publickey := c.Request.FormValue("publickey")
	signature := c.Request.FormValue("signature")
	os.Chdir(".")
	tmpFile, err := os.Create("./static/" + h.Filename + ".pdf")
	if err != nil {
		log.Println("create file error: ", err)
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
	phone, err := dc.userRepo.GetPhoneByPublickey(publickey)
	if err != nil {
		log.Println("get phone error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	convert_signature, err := hex.DecodeString(signature[2:])
	fmt.Printf("% x", convert_signature)
	fmt.Println("\nlength:", len(convert_signature))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	event, err := dc.documentSrv.SaveSignaturalDocument(phone, convert_signature)
	if err != nil {
		log.Println("save signatural error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	doc := model.Document{
		DocID:     uuid.New(),
		Number:    int(event.Numdoc.Int64()),
		Owner:     publickey,
		Name:      h.Filename,
		Type:      "pdf",
		Signature: signature,
		Path:      "static/",
		Public:    true,
		CreateAt:  time.Now(),
		UpdateAt:  time.Time{},
		DeleteAt:  time.Time{},
	}
	err = dc.documentRepo.Create(&doc)
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
		"data":    event,
	})
}

func (dc *DocumentController) GetDocs(c *gin.Context) {
	publickey := c.Param("publickey")
	docs, err := dc.documentSrv.GetDocumentByPublickey(publickey)
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
