package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4/request"
	"net/http"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.AuthorizationHeaderExtractor.ExtractToken(c.Request)
		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		payload, err := ParseToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set(UserId, payload.UserID)
		c.Next()
	}
}
