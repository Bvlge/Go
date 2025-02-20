// controllers/auth_middleware.go
package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware é um exemplo de middleware para autenticação usando JWT.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticação não fornecido"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		_ = tokenString // Remover este comentário após implementar a validação do token

		// Aqui você deve implementar a validação do token.
		// Exemplo:
		// userID, err := validateToken(tokenString)
		// if err != nil {
		//     c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		//     c.Abort()
		//     return
		// }
		// c.Set("userID", userID)

		// Para fins de exemplo, vamos assumir que o token é válido e definir userID = 1.
		c.Set("userID", uint(1))
		c.Next()
	}
}
