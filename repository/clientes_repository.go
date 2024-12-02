package repository

import (
	"fmt"
	"go-api/model"

	"gorm.io/gorm"
)

type ClientesRepository struct {
	connection *gorm.DB
}

// NewProductRepository inicializa o repositório com uma conexão GORM
func NewClientesRepository(connection *gorm.DB) ClientesRepository {
	return ClientesRepository{
		connection: connection,
	}
}

func (r *ClientesRepository) GetClientes(page, maxPerPage int, search string) ([]model.Cliente, int, error) {
	var clienteList []model.Cliente

	// Calcula o offset com base na página e no número máximo de itens por página
	offset := (page - 1) * maxPerPage

	// Base da consulta
	// query := r.connection.Model(&model.Cliente{})

	query := r.connection.Model(&model.Cliente{}).Preload("Consultas") // Carrega as consultas associadas

	// Aplica o filtro de busca se o parâmetro "search" foi fornecido
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	// Conta o número total de clientes com o filtro de busca aplicado
	var totalItems int64
	query.Count(&totalItems)

	// Executa a consulta paginada com limite e offset
	result := query.Limit(maxPerPage).Offset(offset).Find(&clienteList)
	if result.Error != nil {
		fmt.Println(result.Error)
		return []model.Cliente{}, 0, result.Error
	}

	// Calcula o número total de páginas
	totalPages := int((totalItems + int64(maxPerPage) - 1) / int64(maxPerPage))

	return clienteList, totalPages, nil
}

// // CreateProduct cria um novo produto e retorna seu ID
// func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
// 	result := pr.connection.Create(&product)
// 	if result.Error != nil {
// 		fmt.Println(result.Error)
// 		return 0, result.Error
// 	}
// 	return int(product.ID), nil
// }

// // GetProductById obtém um produto pelo ID
// func (pr *ProductRepository) GetProductById(productID int) (*model.Product, error) {
// 	var product model.Product
// 	result := pr.connection.First(&product, productID)
// 	if result.Error != nil {
// 		if result.Error == gorm.ErrRecordNotFound {
// 			return nil, nil
// 		}
// 		fmt.Println(result.Error)
// 		return nil, result.Error
// 	}
// 	return &product, nil
// }

// // RemoveProductById remove um produto pelo ID
// func (pr *ProductRepository) RemoveProductById(product_id int) error {
// 	var product model.Product

// 	// Verifica se o produto existe antes de tentar deletá-lo
// 	result := pr.connection.First(&product, product_id)
// 	if result.Error != nil {
// 		if result.Error == gorm.ErrRecordNotFound {
// 			return nil // Retorna nil para indicar que o produto não existe
// 		}
// 		fmt.Println(result.Error)
// 		return result.Error
// 	}

// 	// Deleta o produto encontrado
// 	deleteResult := pr.connection.Delete(&product)
// 	if deleteResult.Error != nil {
// 		fmt.Println(deleteResult.Error)
// 		return deleteResult.Error
// 	}

// 	return nil
// }
