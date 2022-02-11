package controllers

import (
	"digitalsignature/internal/app/request"
	"digitalsignature/internal/app/service"
	"fmt"

	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func UserRouter(userService service.UserService, r *gin.RouterGroup) {
	uc := UserController{
		userService: userService,
	}
	ar := r.Group("/user")
	ar.POST("/signup", uc.SignUp)
	ar.POST("/signup/:password", uc.SignUpWithPassword)
	ar.POST("/signin", uc.SignIn)
}

func (uc *UserController) SignUp(c *gin.Context) {
	// nhận user info và uccount address từ UI
	userInfo := request.UserInfo{}
	err := c.BindJSON(&userInfo)
	if err != nil {
		fmt.Println(err)
	}
	log.Info("User register request", userInfo)
	err = uc.userService.Create(c, userInfo)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{"status": "true"})
}

func (uc *UserController) SignUpWithPassword(c *gin.Context) {

}

func (uc *UserController) SignIn(c *gin.Context) {

}
