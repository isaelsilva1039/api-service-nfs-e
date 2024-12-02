package routes

import (
	"go-api/controller"
	"go-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	server *gin.Engine,
	productController *controller.ProductController,
	clienteController *controller.ClientesController,
	userController *controller.UserController,
	contratoController *controller.ContratoController,
	empresaController *controller.EmpresaController,
	filialController *controller.FilialController,
	pdvController *controller.PDVController,
	tokenController *controller.TokenController,
) {

	// Rotas protegidas com middleware de autenticação
	auth := server.Group("/api/v1")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/products", productController.GetProducts)
	auth.POST("/product", productController.CreateProduct)
	auth.GET("/product/:productId", productController.GetProductById)
	auth.DELETE("/product/:productId", productController.RemoveProductById)

	// Clientes
	auth.GET("/api/v2/clientes", clienteController.GetClientes)

	// contratos
	auth.GET("/contratos", contratoController.GetAllContratos)
	auth.GET("/contrato/:id", contratoController.GetContratoByID)
	auth.POST("/contratos", contratoController.CreateContrato)
	auth.PUT("/contrato/:id", contratoController.UpdateContrato)
	auth.DELETE("/contrato/:id", contratoController.Remove)

	// empresa
	auth.GET("/empresa", empresaController.GetAllEmpresas)
	auth.GET("/empresa/:id", empresaController.GetEmpresaByID)
	auth.POST("/empresa", empresaController.CreateEmpresa)
	auth.PUT("/empresa/:id", empresaController.UpdateEmpresa)
	auth.DELETE("/empresa/:id", empresaController.RemoveEmpresa)

	// filial
	auth.GET("/filial", filialController.GetAllFilial)
	auth.GET("/filial/:id", filialController.GetFilialByID)
	auth.POST("/filial", filialController.CreateFilial)
	auth.PUT("/filial/:id", filialController.UpdateFilial)
	auth.DELETE("/filial/:id", filialController.RemoveFilial)

	// pdv
	auth.POST("/pdv", pdvController.CreatePDV)
	auth.GET("/pdv", pdvController.GetAllPDVS)

	// pdv
	auth.GET("/pdv/token", tokenController.GetAllTokens)

	// Autenticação
	server.POST("/auth/login", userController.Login)
	server.POST("/auth/register", userController.Register)
}
