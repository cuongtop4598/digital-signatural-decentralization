package controllers

import (
	"digitalsignature/internal/app/model"
	"digitalsignature/internal/app/repository"
	"digitalsignature/internal/app/service/document"
	"digitalsignature/internal/app/utils"
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
	ar.POST("/savesign", dc.Sign)
	ar.POST("/upload", dc.Upload)
	ar.GET("/download/:filename", dc.Download)
	ar.POST("/verify", dc.Verify)
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

	tmpFile, err := os.Create("./static/" + h.Filename + ".pdf")
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
			isPublic := dc.documentRepo.IsPublic(name)
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
	Phone  string   `json:"phone"`
	Digest [32]byte `json:"digest"`
	DocNum int64    `json:"doc_number"`
}

func (dc *DocumentController) Verify(c *gin.Context) {
	verify := VerifyDocRequest{}
	err := c.BindJSON(&verify)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isTrue, err := dc.documentSrv.VerifyDocument(verify.Phone, verify.Digest, big.NewInt(verify.DocNum))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	if isTrue {
		c.JSON(http.StatusOK, gin.H{"message": "true"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "false"})
		return
	}
}

func (dc *DocumentController) Sign(c *gin.Context) {
	err := c.Request.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		log.Println("parse multi part form", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	file, h, err := c.Request.FormFile("doc")
	if err != nil {
		log.Println("get file error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	publickey := c.Request.FormValue("publickey")
	signature := c.Request.FormValue("signature")

	tmpFile, err := os.Create("./static/" + h.Filename + ".pdf")
	if err != nil {
		log.Println("create file error: ", err)
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
		UpdateAt:  time.Time{},
		DeleteAt:  time.Time{},
	}
	err = dc.documentRepo.Create(&doc)
	if err != nil {
		log.Println("create document error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = io.Copy(tmpFile, file)
	if err != nil {
		log.Println("copy file error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	phone, err := dc.userRepo.GetPhoneByPublickey(publickey)
	if err != nil {
		log.Println("get phone error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	event, err := dc.documentSrv.SaveSignaturalDocument(phone, []byte(signature))
	if err != nil {
		log.Println("save signatural error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"doc_info": event})
}
