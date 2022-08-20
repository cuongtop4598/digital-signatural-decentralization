package handlers

import (
	"context"
	"digitalsignature/internal/app/repository"
	"digitalsignature/internal/app/request"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db  *gorm.DB
	log *zap.Logger
	ctx context.Context
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JWTOutput struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

func NewAuthHandler(db *gorm.DB, log *zap.Logger, ctx context.Context) *AuthHandler {
	return &AuthHandler{
		db:  db,
		log: log,
		ctx: ctx,
	}
}

func (handler *AuthHandler) SignInHandler(c *gin.Context) {
	var userLogin request.Login
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userRepo := repository.NewUserRepository(handler.db, handler.log)
	isTrue, userInfo, err := userRepo.CheckLogin(userLogin.Password, userLogin.Phone)
	if err != nil {
		handler.log.Sugar().Error(err)
	}
	if !isTrue {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	expirationTime := time.Now().Add(30000 * time.Hour)
	claims := &Claims{
		Username: userLogin.Phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(
		os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jwtOutput := JWTOutput{
		Token:   tokenString,
		Expires: expirationTime,
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": *userInfo,
		"jwt":  jwtOutput,
	})
}
