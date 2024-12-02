package controller

import (
	"errors"
	dto "go-api/dto/empresa"
	"go-api/usecase"
	"go-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmpresaController struct {
	empresaUsecase usecase.EmpresaUsecase
}

func NewEmpresaController(empresaUsecase *usecase.EmpresaUsecase) *EmpresaController {
	return &EmpresaController{
		empresaUsecase: *empresaUsecase,
	}
}

/** Cria uma nova empresa */
func (c *EmpresaController) CreateEmpresa(ctx *gin.Context) {
	var request dto.CreateEmpresaRequest

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
	err = c.empresaUsecase.CreateEmpresa(request, userID, userType.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Empresa criada com sucesso"})
}

/** Edita uma empresa */
func (c *EmpresaController) UpdateEmpresa(ctx *gin.Context) {
	// Pegue o ID da empresa na URL
	idParam := ctx.Param("id")
	empresaID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Bind dos dados recebidos na requisição
	var request dto.UpdateEmpresaRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chame o caso de uso para atualizar a empresa
	err = c.empresaUsecase.UpdateEmpresa(empresaID, request)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, errors.New("empresa não encontrada")) || errors.Is(err, errors.New("usuário não autorizado a editar esta empresa")) {
			statusCode = http.StatusForbidden
		}
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Empresa atualizada com sucesso"})
}

/** Remove uma empresa */
func (c *EmpresaController) RemoveEmpresa(ctx *gin.Context) {
	// Pegue o ID da empresa na URL
	idParam := ctx.Param("id")
	empresaID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Chame o caso de uso para remover a empresa
	err = c.empresaUsecase.RemoveEmpresa(empresaID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, errors.New("empresa não encontrada")) || errors.Is(err, errors.New("usuário não autorizado a remover esta empresa")) {
			statusCode = http.StatusForbidden
		}
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Empresa removida com sucesso"})
}

/** Obtém todas as empresas */
func (c *EmpresaController) GetAllEmpresas(ctx *gin.Context) {
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
	empresas, total, err := c.empresaUsecase.GetEmpresas(userID.(int), userType.(int), pageInt, pageSizeInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorna as empresas com informações de paginação
	ctx.JSON(http.StatusOK, gin.H{
		"data":        empresas,
		"total":       total,
		"page":        pageInt,
		"page_size":   pageSizeInt,
		"total_pages": (total + int64(pageSizeInt) - 1) / int64(pageSizeInt),
	})
}

/** Obtém uma empresa pelo ID (validando pelo nível do usuário) */
func (c *EmpresaController) GetEmpresaByID(ctx *gin.Context) {
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

	empresa, err := c.empresaUsecase.GetEmpresaByID(id, userID.(int), userType.(int))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Empresa não encontrada ou acesso negado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": empresa})
}
