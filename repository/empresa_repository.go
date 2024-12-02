package repository

import (
	"go-api/model"

	"gorm.io/gorm"
)

type EmpresaRepository struct {
	connection *gorm.DB
}

/** Inicia o repository */
func NewEmpresaRepository(connection *gorm.DB) EmpresaRepository {
	return EmpresaRepository{
		connection: connection,
	}
}

/** pega um contatos por id */
func (r *EmpresaRepository) FindByID(id int) (*model.Empresa, error) {
	var empresa model.Empresa
	if err := r.connection.First(&empresa, id).Error; err != nil {
		return nil, err
	}
	return &empresa, nil
}

/** salva um novo contratos*/
func (r *EmpresaRepository) Save(empresa model.Empresa) error {
	return r.connection.Create(&empresa).Error
}

/** atualiza um contratos*/
func (r *EmpresaRepository) SaveUpdate(empresa *model.Empresa) error {
	return r.connection.Save(empresa).Error
}

/** deleta um contratos*/
func (r *EmpresaRepository) Delete(empresa *model.Empresa) error {
	return r.connection.Delete(empresa).Error
}

/*
  - busca os contratos
    Tipo: 1 pega todos os contratos (perfil adm)
    Tipo: 2 pega só os contratoss que ele cadastrou (perfil parceiro)
*/
func (r *EmpresaRepository) GetAll(userID int, userType int, page int, pageSize int) ([]model.Empresa, int64, error) {
	var empresa []model.Empresa
	var total int64

	offset := (page - 1) * pageSize // Calcula o deslocamento para a paginação

	query := r.connection.Model(model.Empresa{})

	// Filtra com base no tipo de usuário
	if userType == 2 {
		query = query.Where("criado_por = ?", userID)
	}

	// Conta o total de registros antes da paginação
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("id DESC").Offset(offset).Limit(pageSize).Find(&empresa).Error
	if err != nil {
		return nil, 0, err
	}

	return empresa, total, nil
}

/** Pega um contrato por ID com validação de tipo de usuário */
func (r *EmpresaRepository) FindByIDWithUserValidation(id int, userID int, userType int) (*model.Empresa, error) {
	var empresa model.Empresa

	query := r.connection.Model(&model.Empresa{})

	// Se o usuário for do tipo 2 (usuário comum), adicione a validação de propriedade
	if userType == 2 {
		query = query.Where("id = ? AND criado_por = ?", id, userID)
	} else {

		query = query.Where("id = ?", id)
	}

	// Executa a consulta
	if err := query.First(&empresa).Error; err != nil {
		return nil, err
	}

	return &empresa, nil
}
