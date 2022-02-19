package controllers

import (
	"digitalsignature/internal/app/service/document"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type DocumentController struct {
	documentSrv document.DocumentService
}

func DocumentRouter(docService document.DocumentService, r *gin.RouterGroup) {
	dc := DocumentController{
		documentSrv: docService,
	}
	ar := r.Group("/document")
	ar.POST("/upload", dc.Upload)
	ar.GET("/download/:filename", dc.Download)
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
	tmpFile, err := os.Create("./static/" + h.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer tmpFile.Close()

	_, err = io.Copy(tmpFile, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "update load file success"})
}

func (dc *DocumentController) Download(c *gin.Context) {
	name := c.Param("filename")
	c.File("static/" + name)
}

func (dc *DocumentController) Verify(c *gin.Context) {
	name := c.Param("filename")
	c.File("static/" + name)
}
