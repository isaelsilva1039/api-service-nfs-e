package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductCrontroller(usecase usecase.ProductUsecase) *ProductController {
	return &ProductController{
		productUsecase: usecase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {

	products, err := p.productUsecase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {

	var product model.Product

	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertProduct, err := p.productUsecase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertProduct)
}

func (p *ProductController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("productId")

	if id == "" {

		response := model.Response{
			Mensagem: "Id do produto não pode ser null",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	producId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Mensagem: "Id do produto não precisa ser numero",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUsecase.GetProductById(producId)

	if product == nil {

		response := model.Response{
			Mensagem: "Id do produto não precisa ser numero",
		}

		ctx.JSON(http.StatusFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *ProductController) RemoveProductById(ctx *gin.Context) {

	id := ctx.Param("productId")

	if id == "" {

		response := model.Response{
			Mensagem: "Id do produto não pode ser null",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	producId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Mensagem: "Id do produto não precisa ser numero",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = p.productUsecase.RemoveProductById(producId)
	if err != nil {
		if err.Error() == "produto não encontrado" {
			response := model.Response{
				Mensagem: "Produto não encontrado",
			}
			ctx.JSON(http.StatusNotFound, response)
		} else {
			response := model.Response{
				Mensagem: "Erro ao remover produto",
			}
			ctx.JSON(http.StatusInternalServerError, response)
		}
		return
	}

	ctx.JSON(http.StatusOK, producId)
}
