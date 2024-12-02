package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ClientesUsecase struct {
	repository repository.ClientesRepository
}

func NewclientesUsecase(repo repository.ClientesRepository) ClientesUsecase {
	return ClientesUsecase{
		repository: repo,
	}
}

// GetClientes busca a lista de clientes com paginação
func (u *ClientesUsecase) GetClientes(page, maxPerPage int, search string) ([]model.Cliente, int, error) {
	return u.repository.GetClientes(page, maxPerPage, search)
}

// func (pu *ClientesUsecase) CreateProduct(product model.Product) (model.Product, error) {
// 	productId, err := pu.repository.CreateProduct(product)

// 	if err != nil {
// 		return model.Product{}, err
// 	}

// 	product.ID = productId

// 	return product, nil
// }

// func (pu *ProductUsecase) GetProductById(product_id int) (*model.Product, error) {

// 	product, err := pu.repository.GetProductById(product_id)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return product, nil
// }

// func (pu *ProductUsecase) RemoveProductById(product_id int) error {
// 	// Apenas verifica o erro ao chamar a função de remoção
// 	err := pu.repository.RemoveProductById(product_id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
