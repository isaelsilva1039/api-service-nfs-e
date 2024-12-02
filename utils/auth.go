package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// GetUserIDFromContext extrai o userID do contexto
func GetUserIDFromContext(ctx *gin.Context) (int, error) {
	userID, exists := ctx.Get("userID") // Recupera o userID do contexto
	if !exists {
		return 0, errors.New("usuário não autenticado")
	}

	// Converte o userID para int, se necessário
	id, ok := userID.(int)
	if !ok {
		return 0, errors.New("ID do usuário no contexto está inválido")
	}

	return id, nil
}
