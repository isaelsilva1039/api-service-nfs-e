package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClientesController struct {
	ClientesUsecase usecase.ClientesUsecase
}

func NewClientesCrontroller(usecase usecase.ClientesUsecase) *ClientesController {
	return &ClientesController{
		ClientesUsecase: usecase,
	}
}

func (c *ClientesController) GetClientes(ctx *gin.Context) {
	page := 1
	maxPerPage := 20

	if p := ctx.Query("page"); p != "" {
		if val, err := strconv.Atoi(p); err == nil && val > 0 {
			page = val
		}
	}
	if m := ctx.Query("max_per_page"); m != "" {
		if val, err := strconv.Atoi(m); err == nil && val > 0 {
			maxPerPage = val
		}
	}

	search := ctx.Query("search")

	clientes, totalPages, err := c.ClientesUsecase.GetClientes(page, maxPerPage, search)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := model.ResponseCliente{
		Data: clientes,
		Meta: model.PaginationMeta{
			MaxPerPage: maxPerPage,
			Page:       page,
			TotalItems: totalPages * maxPerPage,
			TotalPages: totalPages,
		},
	}

	ctx.JSON(http.StatusOK, response)
}

// func (p *productController) CreateProduct(ctx *gin.Context) {

// 	var product model.Product

// 	err := ctx.BindJSON(&product)

// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	insertProduct, err := p.productUsecase.CreateProduct(product)

// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, err)
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, insertProduct)
// }

// func (p *productController) GetProductById(ctx *gin.Context) {

// 	id := ctx.Param("productId")

// 	if id == "" {

// 		response := model.Response{
// 			Mensagem: "Id do produto não pode ser null",
// 		}

// 		ctx.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	producId, err := strconv.Atoi(id)

// 	if err != nil {
// 		response := model.Response{
// 			Mensagem: "Id do produto não precisa ser numero",
// 		}

// 		ctx.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	product, err := p.productUsecase.GetProductById(producId)

// 	if product == nil {

// 		response := model.Response{
// 			Mensagem: "Id do produto não precisa ser numero",
// 		}

// 		ctx.JSON(http.StatusFound, response)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, product)
// }

// func (p *productController) RemoveProductById(ctx *gin.Context) {

// 	id := ctx.Param("productId")

// 	if id == "" {

// 		response := model.Response{
// 			Mensagem: "Id do produto não pode ser null",
// 		}

// 		ctx.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	producId, err := strconv.Atoi(id)

// 	if err != nil {
// 		response := model.Response{
// 			Mensagem: "Id do produto não precisa ser numero",
// 		}

// 		ctx.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	err = p.productUsecase.RemoveProductById(producId)
// 	if err != nil {
// 		if err.Error() == "produto não encontrado" {
// 			response := model.Response{
// 				Mensagem: "Produto não encontrado",
// 			}
// 			ctx.JSON(http.StatusNotFound, response)
// 		} else {
// 			response := model.Response{
// 				Mensagem: "Erro ao remover produto",
// 			}
// 			ctx.JSON(http.StatusInternalServerError, response)
// 		}
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, producId)
// }
