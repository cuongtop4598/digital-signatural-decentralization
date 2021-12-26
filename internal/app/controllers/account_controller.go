package controllers

import (
	"digitalsignature/internal/app/service"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountService service.AccountService
}

func AccountRouter(accountService service.AccountService, r *gin.RouterGroup) {
	ac := AccountController{
		accountService: accountService,
	}
	ar := r.Group("/account")
	ar.GET("/signup", ac.SignUp)
	ar.POST("/signin", ac.SignIn)
}

func (ac *AccountController) SignUp(c *gin.Context) {
	key, err := ac.accountService.CreateNewAccount(c)
	if err != nil {
		c.JSON(400, gin.H{"message": "Create account ERROR :", "error": err.Error()})
	}
	c.JSON(200, key)
}

func (ac *AccountController) SignIn(c *gin.Context) {

}
