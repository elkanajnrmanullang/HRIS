package middleware

import (
	"net/http"
	"strings"

	"celestara.com/hris-api/internal/shared/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("super-secret-key-celestara")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, response.Error("Authorization header required", nil))
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, response.Error("Invalid authorization format", nil))
			c.Abort()
			return
		}

		tokenString := parts[1]
		
		// Validasi token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, response.Error("Invalid or expired token", nil))
			c.Abort()
			return
		}

		// Simpan klaim ke context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["user_id"])
			c.Set("role", claims["role"])
		}

		c.Next()
	}
}