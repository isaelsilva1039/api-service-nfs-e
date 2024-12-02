package controller

import (
	"errors"
	dto "go-api/dto/contrato"
	"go-api/usecase"
	"go-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContratoController struct {
	contratoUsecase usecase.ContratoUsecase
}

func NewContratoController(contratoUsecase *usecase.ContratoUsecase) *ContratoController {
	return &ContratoController{
		contratoUsecase: *contratoUsecase,
	}
}

/**  Cria um novo contratos */
func (c *ContratoController) CreateContrato(ctx *gin.Context) {
	var request dto.CreateContratoRequest // Use a estrutura definida no DTO

	if err := ctx.ShouldBindJSON(&request); err != nil { // Faz o bind dos dados recebidos na requisição
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Pegue o usuário logado
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	// Chame o caso de uso
	err = c.contratoUsecase.CreateContrato(request, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Contrato criado com sucesso"})
}

/**  edita um contratos */
func (c *ContratoController) UpdateContrato(ctx *gin.Context) {
	// Pegue o ID do contrato na URL
	idParam := ctx.Param("id")
	contratoID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Pegue o usuário logado
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	// Bind dos dados recebidos na requisição
	var request dto.UpdateContratoRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chame o caso de uso para atualizar o contrato
	err = c.contratoUsecase.UpdateContrato(contratoID, request, userID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, errors.New("contrato não encontrado")) || errors.Is(err, errors.New("usuário não autorizado a editar este contrato")) {
			statusCode = http.StatusForbidden
		}
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Contrato atualizado com sucesso"})
}

/**  remove um contratos */
func (c *ContratoController) Remove(ctx *gin.Context) {

	// Pegue o ID do contrato na URL
	idParam := ctx.Param("id")
	contratoID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Chame o caso de uso para atualizar o contrato
	err = c.contratoUsecase.Remove(contratoID)
	if err != nil {

		statusCode := http.StatusInternalServerError
		if errors.Is(err, errors.New("contrato não encontrado")) {
			statusCode = http.StatusForbidden
		}

		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Contrato removido com sucesso"})
}

/**  obtem contratos */
func (c *ContratoController) GetAllContratos(ctx *gin.Context) {
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
	contratos, total, err := c.contratoUsecase.GetContratos(userID.(int), userType.(int), pageInt, pageSizeInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorna os contratos com informações de paginação
	ctx.JSON(http.StatusOK, gin.H{
		"data":        contratos,
		"total":       total,
		"page":        pageInt,
		"page_size":   pageSizeInt,
		"total_pages": (total + int64(pageSizeInt) - 1) / int64(pageSizeInt),
	})
}

/**  obtem 1 contatos contratos (validando pelo nivel do ususario) */
func (c *ContratoController) GetContratoByID(ctx *gin.Context) {
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

	contrato, err := c.contratoUsecase.GetContratoByID(id, userID.(int), userType.(int))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Contrato não encontrado ou acesso negado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": contrato})
}
