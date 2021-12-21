package controllers

import (
	"digitalsignature/internal/app/service"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountService *service.AccountService
}

func AccountRouter(accountService *service.AccountService, r *gin.RouterGroup) {
	ac := AccountController{
		accountService: accountService,
	}
	ar := r.Group("/account")
	ar.POST("/account/signup", ac.SignUp)
	ar.POST("/account/signin", ac.SignIn)
}

func (ac *AccountController) SignUp(c *gin.Context) {

}

func (ac *AccountController) SignIn(c *gin.Context) {

}
