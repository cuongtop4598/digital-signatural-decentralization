package controllers

import (
	"digitalsignature/internal/app/request"
	"digitalsignature/internal/app/service"
	"net/http"

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
	ar.GET("/info/:pubkey", uc.GetInfo)
	ar.POST("/verify", uc.Verify)
	ar.POST("/signin", uc.SignIn)
}

func (uc *UserController) SignUp(c *gin.Context) {
	// nhận user info và uccount address từ UI
	userInfo := request.UserInfo{}
	err := c.BindJSON(&userInfo)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "false"})
		return
	}
	if userInfo.Phone == "" || userInfo.PublicKey == "" || userInfo.Password == "" {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "false"})
		return
	}
	log.Info("User register request", userInfo)
	err = uc.userService.Create(c, userInfo)
	if err != nil {
		log.Error(err.Error())
		c.JSON(404, gin.H{"status": "false"})
		return
	}
	c.SetCookie("publickey", userInfo.PublicKey, 10000000, "", "", false, false)
	c.SetCookie("phone", userInfo.Phone, 10000000, "", "", false, false)
	c.JSON(200, gin.H{"status": "true"})
}

func (uc *UserController) SignIn(c *gin.Context) {
	loginInfo := request.Login{}
	err := c.BindJSON(&loginInfo)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "false"})
		return
	}

	// bad security, i will change this verify later
	if loginInfo.Phone == "" || loginInfo.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false"})
		return
	}

	isLog, userInfo, err := uc.userService.Login(loginInfo)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	if isLog {
		c.SetCookie("publickey", userInfo.PublicKey, 10000000, "", "", false, false)
		c.SetCookie("phone", userInfo.Phone, 10000000, "", "", false, false)
		c.JSON(http.StatusAccepted, gin.H{"login": true})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"login": false})
	}
}

func (uc *UserController) GetInfo(c *gin.Context) {
	pubkey := c.Param("pubkey")
	response, err := uc.userService.GetUserInfo(c, pubkey)
	if err != nil {
		log.Error(err.Error())
		c.JSON(404, gin.H{"status": "false"})
		return
	}
	c.JSON(200, response)
}

func (uc *UserController) Verify(c *gin.Context) {
	userInfo := request.UserInfo{}
	err := c.BindJSON(&userInfo)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "can't decode user info request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "user information is correct"})
}
