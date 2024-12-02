package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token malformado"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, gin.Error{Err: http.ErrAbortHandler}
			}
			return []byte("secrethash"), nil // Use a mesma chave usada para gerar o token
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Extraia as informações do token
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Exemplo: Extraindo o `userID` do payload do JWT
			userID, ok := claims["userID"].(float64) // `float64` é o padrão para números no JWT
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "userID não encontrado no token"})
				c.Abort()
				return
			}

			// Extraia o `userType`
			userType, ok := claims["userType"].(float64)
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Tipo de usuário não encontrado no token"})
				c.Abort()
				return
			}

			c.Set("userType", int(userType))

			c.Set("userID", int(userID)) // Converte para `int` antes de salvar no contexto
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Falha ao processar token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
