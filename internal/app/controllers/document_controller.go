package controllers

import (
	"digitalsignature/internal/app/service/document"

	"github.com/gin-gonic/gin"
)

type DocumentController struct {
	documentService document.DocumentService
}

func DocumentRouter(documentService document.DocumentService, r *gin.RouterGroup) {
	dc := DocumentController{
		documentService: documentService,
	}
	dr := r.Group("/document")
	dr.POST("/upload/public", dc.UploadPublic)
	dr.POST("/upload/private")
	dr.POST("/download", dc.Download)
	dr.POST("/verify", dc.Verify)
}

// Upload document using IPFS
func (d *DocumentController) UploadPublic(c *gin.Context) {

}

// Upload document and save to server
func (d *DocumentController) UploadPrivate(c *gin.Context) {

}

// Upload
func (d *DocumentController) Verify(c *gin.Context) {

}

func (d *DocumentController) Download(c *gin.Context) {

}
