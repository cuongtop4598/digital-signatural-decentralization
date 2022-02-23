package controllers

import (
	"digitalsignature/internal/app/model"
	"digitalsignature/internal/app/repository"
	"digitalsignature/internal/app/request"
	"digitalsignature/internal/app/service/document"
	"digitalsignature/internal/app/utils"
	"fmt"
	"io"
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
}

func DocumentRouter(docService document.DocumentService, documentRepo repository.DocumentRepository, r *gin.RouterGroup) {
	dc := DocumentController{
		documentSrv:  docService,
		documentRepo: documentRepo,
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

	tmpFile, err := os.Create("./static/" + h.Filename)
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
		UpdateAt:  time.Time{},
		DeleteAt:  time.Time{},
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
	}
}

func (dc *DocumentController) Sign(c *gin.Context) {
	save := request.SaveSignatural{}
	err := c.BindJSON(&save)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	events := dc.documentSrv.SaveSignaturalDocument(save.Phone, save.Signatural)
	c.JSON(http.StatusOK, events)
}
