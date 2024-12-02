package controller

import (
	"errors"
	dto "go-api/dto/filial"
	"go-api/usecase"
	"go-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/** struct */
type FilialController struct {
	filialUsecase usecase.FilialUsecase
}

/** construtor inicial */
func NewFilialController(filialUsecase *usecase.FilialUsecase) *FilialController {
	return &FilialController{
		filialUsecase: *filialUsecase,
	}
}

/** Cria uma nova filial */
func (c *FilialController) CreateFilial(ctx *gin.Context) {
	var request dto.CreateFilialRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Pegue o usuário logado e tipo
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	userType, exists := ctx.Get("userType")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Tipo de usuário não encontrado"})
		return
	}

	// Chame o caso de uso
	err = c.filialUsecase.CreateFilial(request, userID, userType.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Filial criada com sucesso"})
}

/** Edita uma filial */
func (c *FilialController) UpdateFilial(ctx *gin.Context) {
	// Pegue o ID da filial na URL
	idParam := ctx.Param("id")
	filialID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Bind dos dados recebidos na requisição
	var request dto.UpdateFilialRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chame o caso de uso para atualizar a filial
	err = c.filialUsecase.UpdateFilial(filialID, request)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, errors.New("filial não encontrada")) || errors.Is(err, errors.New("usuário não autorizado a editar esta filial")) {
			statusCode = http.StatusForbidden
		}
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Filial atualizada com sucesso"})
}

/** Remove uma filial */
func (c *FilialController) RemoveFilial(ctx *gin.Context) {
	// Pegue o ID da filial na URL
	idParam := ctx.Param("id")
	filialID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Chame o caso de uso para remover a filial
	err = c.filialUsecase.RemoveFilial(filialID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, errors.New("filial não encontrada")) || errors.Is(err, errors.New("usuário não autorizado a remover esta filial")) {
			statusCode = http.StatusForbidden
		}
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Filial removida com sucesso"})
}

/** Obtém todas as filial */
func (c *FilialController) GetAllFilial(ctx *gin.Context) {
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
	filiais, total, err := c.filialUsecase.GetFilial(userID.(int), userType.(int), pageInt, pageSizeInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorna as filial com informações de paginação
	ctx.JSON(http.StatusOK, gin.H{
		"data":        filiais,
		"total":       total,
		"page":        pageInt,
		"page_size":   pageSizeInt,
		"total_pages": (total + int64(pageSizeInt) - 1) / int64(pageSizeInt),
	})
}

/** Obtém uma filial pelo ID (validando pelo nível do usuário) */
func (c *FilialController) GetFilialByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

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

	filial, err := c.filialUsecase.GetFilialByID(id, userID.(int), userType.(int))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Filial não encontrada ou acesso negado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": filial})
}
