package controllers

import (
	"digitalsignature/internal/app/model"
	"digitalsignature/internal/app/repository"
	"digitalsignature/internal/app/service/document"
	"digitalsignature/internal/app/utils"
	"fmt"
	"io"
	"net/http"
	"os"

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
	ar.POST("/upload", dc.Upload)
	ar.GET("/download/:filename", dc.Download)
	ar.POST("/verify/:filename", dc.Verify)
}

func (dc *DocumentController) Upload(c *gin.Context) {
	userID := c.Param("user_id")
	public := c.Param("public")
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
	tmpFile, err := os.Create("./static/" + h.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer tmpFile.Close()

	var doc model.Document

	if public == "true" {
		doc = model.Document{
			Owner:  uuid.MustParse(userID),
			Name:   h.Filename,
			Type:   "pdf",
			Path:   "static/",
			Public: true,
		}
	} else {
		doc = model.Document{
			Owner:  uuid.MustParse(userID),
			Name:   h.Filename,
			Type:   "pdf",
			Path:   "static/",
			Public: false,
		}
	}
	_, err = io.Copy(tmpFile, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = dc.documentRepo.Create(&doc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "update load file success"})
}

func (dc *DocumentController) Download(c *gin.Context) {
	name := c.Param("filename")
	files := utils.SearchFileInPath("static/")
	fmt.Println(files)
	for _, file := range files {
		if file == ("static/" + name) {
			c.File(file)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "file not found"})
}

func (dc *DocumentController) Verify(c *gin.Context) {
}
