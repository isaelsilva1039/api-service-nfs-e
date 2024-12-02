package controller

import (
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TokenController struct {
	tokenUsecase *usecase.TokenUsecase
}

func NewTokenController(tokenUsecase *usecase.TokenUsecase) *TokenController {
	return &TokenController{tokenUsecase: tokenUsecase}
}

/** Obtém todas as empresas */
func (c *TokenController) GetAllTokens(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	userType, exists := ctx.Get("userType")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Tipo de usuário não encontrado"})
		return
	}

	// Obtenha os parâmetros de paginação da URL
	page := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("page_size", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetro 'page' inválido"})
		return
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetro 'page_size' inválido"})
		return
	}

	// Chame o caso de uso com paginação
	tokens, total, err := c.tokenUsecase.GetAllTokens(userID.(int), userType.(int), pageInt, pageSizeInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorna as empresas com informações de paginação
	ctx.JSON(http.StatusOK, gin.H{
		"data":        tokens,
		"total":       total,
		"page":        pageInt,
		"page_size":   pageSizeInt,
		"total_pages": (total + int64(pageSizeInt) - 1) / int64(pageSizeInt),
	})
}
