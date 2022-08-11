package controllers

import (
	"digitalsignature/internal/app/request"
	"digitalsignature/internal/app/service"
	"net/http"

	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func UserRouter(userService *service.UserService, r *gin.RouterGroup) {
	uc := UserController{
		userService: userService,
	}
	ar := r.Group("/user")
	ar.POST("/register", uc.Register)
	ar.GET("/info/:phone", uc.GetInfo)
	ar.POST("/verify", uc.Verify)
	ar.POST("/login", uc.Login)
}

func (uc *UserController) Register(c *gin.Context) {
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
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Fail",
			"data":    "",
		})
		return
	}
	userInfo.SantisizePassword()
	c.JSON(200, gin.H{
		"code":    200,
		"message": "OK",
		"data":    userInfo,
	})
}

func (uc *UserController) Login(c *gin.Context) {
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	if isLog {
		userInfo.SantisizePassword()
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "OK",
			"data":    userInfo,
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"message": "Forbidden",
			"data":    "",
		})
	}
}

func (uc *UserController) GetInfo(c *gin.Context) {
	phone := c.Param("phone")
	response, err := uc.userService.GetUserInfo(c, phone)
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "can't decode user info request"})
		return
	}
	isTrue, err := uc.userService.VerifyUser(userInfo)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if isTrue {
		c.JSON(http.StatusOK, gin.H{"message": "user information is correct"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "user information is wrong"})
	}
}
