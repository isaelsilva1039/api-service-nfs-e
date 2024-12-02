package repository

import (
	"go-api/model"

	"gorm.io/gorm"
)

type ContratoRepository struct {
	connection *gorm.DB
}

/** Inicia o repository */
func NewContratoRepository(connection *gorm.DB) ContratoRepository {
	return ContratoRepository{
		connection: connection,
	}
}

/** pega um contatos por id */
func (r *ContratoRepository) FindByID(id int) (*model.Contrato, error) {
	var contrato model.Contrato
	if err := r.connection.First(&contrato, id).Error; err != nil {
		return nil, err
	}
	return &contrato, nil
}

/** salva um novo contratos*/
func (r *ContratoRepository) Save(contrato model.Contrato) error {
	return r.connection.Create(&contrato).Error
}

/** atualiza um contratos*/
func (r *ContratoRepository) SaveUpdate(contrato *model.Contrato) error {
	return r.connection.Save(contrato).Error
}

/** deleta um contratos*/
func (r *ContratoRepository) Delete(contrato *model.Contrato) error {
	return r.connection.Delete(contrato).Error
}

/*
  - busca os contratos
    Tipo: 1 pega todos os contratos (perfil adm)
    Tipo: 2 pega só os contratoss que ele cadastrou (perfil parceiro)
*/
func (r *ContratoRepository) GetAll(userID int, userType int, page int, pageSize int) ([]model.Contrato, int64, error) {
	var contratos []model.Contrato
	var total int64

	offset := (page - 1) * pageSize // Calcula o deslocamento para a paginação

	query := r.connection.Model(model.Contrato{})

	// Filtra com base no tipo de usuário
	if userType == 2 {
		query = query.Where("criado_por = ?", userID)
	}

	// Conta o total de registros antes da paginação
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Aplica paginação e busca os registros
	err = query.Order("id DESC").Offset(offset).Limit(pageSize).Find(&contratos).Error
	if err != nil {
		return nil, 0, err
	}

	return contratos, total, nil
}

/** Pega um contrato por ID com validação de tipo de usuário */
func (r *ContratoRepository) FindByIDWithUserValidation(id int, userID int, userType int) (*model.Contrato, error) {
	var contrato model.Contrato

	query := r.connection.Model(&model.Contrato{})

	// Se o usuário for do tipo 2 (usuário comum), adicione a validação de propriedade
	if userType == 2 {
		query = query.Where("id = ? AND criado_por = ?", id, userID)
	} else {
		// Caso contrário, apenas busque pelo ID (usuário admin)
		query = query.Where("id = ?", id)
	}

	// Executa a consulta
	if err := query.First(&contrato).Error; err != nil {
		return nil, err
	}

	return &contrato, nil
}
