package controllers

import (
	"digitalsignature/internal/app/errors"
	"digitalsignature/internal/app/handlers"
	"digitalsignature/internal/app/request"
	"digitalsignature/internal/app/service"
	"net/http"

	"gorm.io/gorm"

	"go.uber.org/zap"

	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	db          *gorm.DB
	log         *zap.Logger
	userService *service.UserService
}

func UserRouter(db *gorm.DB, log *zap.Logger, userService *service.UserService, r *gin.RouterGroup) {
	uc := UserController{
		db:          db,
		log:         log,
		userService: userService,
	}
	ar := r.Group("/user")
	ar.POST("/register", uc.Register)
	ar.POST("/confirm", uc.Confirm)
	ar.GET("/profile/:phone", uc.GetProfile)
	ar.POST("/verify", uc.Verify)
	ar.POST("/login", uc.Login)
	ar.GET("/confirmed", uc.GetUserConfirmed)
	ar.GET("/unconfirmed", uc.GetUserUnConfirmed)
}

func (uc *UserController) Confirm(c *gin.Context) {
	userInfo := request.UserConfirm{}
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		uc.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "false"})
		return
	}
	uc.log.Sugar().Info(userInfo)
	if len(userInfo.Phone) <= 0 || len(userInfo.PublicKey) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false"})
		return
	}
	err = uc.userService.Confirm(c, userInfo)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, err)
	}
	c.JSON(http.StatusOK, gin.H{"status": true})
}

func (uc *UserController) Register(c *gin.Context) {
	userInfo := request.UserInfo{}
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "false"})
		return
	}
	uc.log.Sugar().Info(userInfo)
	if userInfo.Phone == "" || userInfo.PublicKey == "" || userInfo.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false"})
		return
	}

	err = uc.userService.Register(c, userInfo)
	switch err {
	case nil:
		break
	case errors.UserIsExisted:
		c.JSON(http.StatusOK, errors.UserIsExisted)
		return
	default:
		c.JSON(http.StatusNotAcceptable, err)
		return
	}
	userInfo.SantisizePassword()
	c.JSON(200, gin.H{
		"code": 200,
		"user": userInfo,
	})
}

func (uc *UserController) Login(c *gin.Context) {
	authHandler := handlers.NewAuthHandler(uc.db, uc.log, c)
	authHandler.SignInHandler(c)
}

func (uc *UserController) GetProfile(c *gin.Context) {
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

func (uc *UserController) GetUserConfirmed(c *gin.Context) {
	users := uc.userService.GetUserConfirmed()
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetUserUnConfirmed(c *gin.Context) {
	users := uc.userService.GetUserUnConfirmed()
	c.JSON(http.StatusOK, users)
}
