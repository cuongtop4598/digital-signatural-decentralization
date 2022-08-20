package middleware

import (
	"digitalsignature/internal/app/handlers"
	"encoding/json"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtOut := handlers.JWTOutput{}
		tokenValue := c.GetHeader("Authorization")
		claims := &handlers.Claims{}
		json.Unmarshal([]byte(tokenValue), &jwtOut)
		tkn, err := jwt.ParseWithClaims(jwtOut.Token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		if tkn == nil || !tkn.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}
