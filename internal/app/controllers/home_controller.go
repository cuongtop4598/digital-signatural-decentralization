package controllers

import (
	"github.com/gin-gonic/gin"
)

type HomeController struct {
}

func HomeRouter(r *gin.Engine) {
	controller := HomeController{}
	r.GET("/", controller.Index)
	r.GET("/healthz", controller.Healthz)
}

func (ac *HomeController) Index(c *gin.Context) {
	c.JSON(200, "Welcome to digital signature service! A Product of MTA development")
}

func (ac *HomeController) Healthz(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  200,
		"healthz": "ok",
	})
}
