package repository

import (
	"fmt"
	"go-api/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	connection *gorm.DB
}

// NewProductRepository inicializa o repositório com uma conexão GORM
func NewProductRepository(connection *gorm.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

// GetProducts obtém todos os produtos
func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	var productList []model.Product
	result := pr.connection.Find(&productList)
	if result.Error != nil {
		fmt.Println(result.Error)
		return []model.Product{}, result.Error
	}
	return productList, nil
}

// CreateProduct cria um novo produto e retorna seu ID
func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	result := pr.connection.Create(&product)
	if result.Error != nil {
		fmt.Println(result.Error)
		return 0, result.Error
	}
	return int(product.ID), nil
}

// GetProductById obtém um produto pelo ID
func (pr *ProductRepository) GetProductById(productID int) (*model.Product, error) {
	var product model.Product
	result := pr.connection.First(&product, productID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		fmt.Println(result.Error)
		return nil, result.Error
	}
	return &product, nil
}

// RemoveProductById remove um produto pelo ID
func (pr *ProductRepository) RemoveProductById(product_id int) error {
	var product model.Product

	// Verifica se o produto existe antes de tentar deletá-lo
	result := pr.connection.First(&product, product_id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil // Retorna nil para indicar que o produto não existe
		}
		fmt.Println(result.Error)
		return result.Error
	}

	// Deleta o produto encontrado
	deleteResult := pr.connection.Delete(&product)
	if deleteResult.Error != nil {
		fmt.Println(deleteResult.Error)
		return deleteResult.Error
	}

	return nil
}
