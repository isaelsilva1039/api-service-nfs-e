package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/routes"
	"go-api/usecase"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// Configurações de CORS
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Permitir todas as origens
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60, // Cache de 12 horas
	}))

	// Conexão com o banco de dados
	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	/** Camada repository */
	productRepository := repository.NewProductRepository(dbConnection)   // Repositório de produtos
	clienteRepository := repository.NewClientesRepository(dbConnection)  // Repositório de clientes
	userRepository := repository.NewUserRepository(dbConnection)         // Repositório de usuários
	contratoRepository := repository.NewContratoRepository(dbConnection) // Repositório de contratos
	empresaRepository := repository.NewEmpresaRepository(dbConnection)   // Repositório de empresas
	filialRepository := repository.NewFilialRepository(dbConnection)     // Repositório de filial
	pdvRepository := repository.NewPDVRepository(dbConnection)           // Repositório de pdv
	tokensRepository := repository.NewTokenRepository(dbConnection)      // Repositório de token

	/** Camada casos de uso */
	productUsecase := usecase.NewproductUsecase(productRepository)    // Caso de uso de produtos
	clienteUsecase := usecase.NewclientesUsecase(clienteRepository)   // Caso de uso de clientes
	userUseCase := usecase.NewUserUseCase(userRepository)             // Caso de uso de usuários
	contratoUsecase := usecase.NewContratoUsecase(contratoRepository) // Caso de uso de contratos
	empresaUsecase := usecase.NewEmpresaUsecase(empresaRepository)    // Caso de uso de empresas
	filialUsecase := usecase.NewFilialUsecase(filialRepository)       // Caso de uso de filial
	pdvUsecase := usecase.NewPDVUsecase(pdvRepository)                // Caso de uso de pdv
	tokenUsecase := usecase.NewTokenUsecase(tokensRepository)         // Caso de uso de tokens

	/** Camada controller */
	productController := controller.NewProductCrontroller(productUsecase)   // Controller de produtos
	clienteController := controller.NewClientesCrontroller(clienteUsecase)  // Controller de clientes
	userController := controller.NewUserController(userUseCase)             // Controller de usuários
	contratoController := controller.NewContratoController(contratoUsecase) // Controller de contratos
	empresaController := controller.NewEmpresaController(empresaUsecase)    // Controller de empresas
	filialController := controller.NewFilialController(filialUsecase)       // Controller de filiasi
	pdvController := controller.NewPDVController(pdvUsecase)                // Controller de pdvs
	tokenController := controller.NewTokenController(tokenUsecase)          // Controller de tokens

	/** Configurar rotas */
	routes.SetupRoutes(
		server,
		productController,
		clienteController,
		userController,
		contratoController,
		empresaController,
		filialController,
		pdvController,
		tokenController,
	)

	// Iniciar o servidor
	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
