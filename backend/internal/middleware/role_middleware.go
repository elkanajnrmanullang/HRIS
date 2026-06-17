package middleware

import (
	"net/http"

	"celestara.com/hris-api/internal/shared/response"
	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		
		if !exists || userRole != requiredRole {
			c.JSON(http.StatusForbidden, response.Error("Access denied: insufficient permission", nil))
			c.Abort()
			return
		}

		c.Next()
	}
}