package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware es un middleware para la autenticación
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Middleware de autenticación invocado")
		token := c.GetHeader("Authorization")
		if token != "mi_token_secreto" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}
