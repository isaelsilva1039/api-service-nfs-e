package controller

import (
	dto "go-api/dto/pdv"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PDVController struct {
	pdvUsecase *usecase.PDVUsecase
}

func NewPDVController(pdvUsecase *usecase.PDVUsecase) *PDVController {
	return &PDVController{pdvUsecase: pdvUsecase}
}

func (c *PDVController) CreatePDV(ctx *gin.Context) {
	var req dto.CreatePDVRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	userID, exists := ctx.Get("userID")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	response, err := c.pdvUsecase.CreatePDV(req, userID.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

/** Obtém todas as empresas */
func (c *PDVController) GetAllPDVS(ctx *gin.Context) {
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
	pdvs, total, err := c.pdvUsecase.GetAllPdvs(userID.(int), userType.(int), pageInt, pageSizeInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorna as empresas com informações de paginação
	ctx.JSON(http.StatusOK, gin.H{
		"data":        pdvs,
		"total":       total,
		"page":        pageInt,
		"page_size":   pageSizeInt,
		"total_pages": (total + int64(pageSizeInt) - 1) / int64(pageSizeInt),
	})
}
