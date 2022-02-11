package controllers

import (
	"digitalsignature/internal/app/request"
	"digitalsignature/internal/app/service"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountService service.UserService
}

func AccountRouter(accountService service.UserService, r *gin.RouterGroup) {
	ac := AccountController{
		accountService: accountService,
	}
	ar := r.Group("/account")
	ar.POST("/signup", ac.SignUp)
	ar.POST("/signup/:password", ac.SignUpWithPassword)
	ar.POST("/signin", ac.SignIn)
}

func (ac *AccountController) SignUp(c *gin.Context) {
	// nhận user info và account address từ UI
	userInfo := request.UserInfo{}
	err := c.BindJSON(userInfo)
	if err != nil {
		c.JSON(400, gin.H{"status": "false", "error": err.Error()})
	}

	c.JSON(200, gin.H{"status": "true"})
}

func (ac *AccountController) SignUpWithPassword(c *gin.Context) {

}

func (ac *AccountController) SignIn(c *gin.Context) {

}
