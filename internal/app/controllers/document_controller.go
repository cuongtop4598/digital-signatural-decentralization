package controllers

import (
	"digitalsignature/internal/app/request"
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
	dr.POST("/sign", dc.SignDocument)
	dr.POST("/verify", dc.VerifyDocument)
}

func (d *DocumentController) SignDocument(c *gin.Context) {
	signDocInfo := request.SignDocumentRequest{}
	err := c.BindJSON(signDocInfo)
	if err != nil {
		c.JSON(200, gin.H{"message": "your request is wrong format"})
	} else {
		docID := d.documentService.SaveDocument(signDocInfo.UserID, "")
		c.JSON(200, gin.H{"DocID": docID, "State": "Success"})
	}

}

func (d *DocumentController) VerifyDocument(c *gin.Context) {

}
