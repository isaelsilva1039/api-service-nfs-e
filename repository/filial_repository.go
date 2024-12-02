package repository

import (
	"errors"
	"go-api/model"

	"gorm.io/gorm"
)

type FilialRepository struct {
	connection *gorm.DB
}

/** Inicia o repository */
func NewFilialRepository(connection *gorm.DB) FilialRepository {
	return FilialRepository{
		connection: connection,
	}
}

/** salva um novo filial*/
func (r *FilialRepository) Save(filial model.Filial) error {
	return r.connection.Create(&filial).Error
}

/** atualiza uma filial */
func (r *FilialRepository) SaveUpdate(filial *model.Filial) error {
	// Verifica se o CNPJ já existe em outro registro
	var existingFilial model.Filial
	err := r.connection.Where("cnpj = ? AND id != ?", filial.CNPJ, filial.ID).First(&existingFilial).Error
	if err == nil {
		return errors.New("CNPJ já está sendo usado por outra filial")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// Retorna erro inesperado, exceto quando não encontrado
		return err
	}

	// Salva as alterações no registro
	return r.connection.Save(filial).Error
}

/** deleta um filial*/
func (r *FilialRepository) Delete(filial *model.Filial) error {
	return r.connection.Delete(filial).Error
}

/** pega um filial por id */
func (r *FilialRepository) FindByID(id int) (*model.Filial, error) {
	var filial model.Filial
	if err := r.connection.Preload("Empresa").First(&filial, id).Error; err != nil {
		return nil, err
	}
	return &filial, nil
}

/*
  - busca os filial
    Tipo: 1 pega todos os filial (perfil adm)
    Tipo: 2 pega só os filials que ele cadastrou (perfil parceiro)
*/
func (r *FilialRepository) GetAll(userID int, userType int, page int, pageSize int) ([]model.Filial, int64, error) {
	var filial []model.Filial
	var total int64

	offset := (page - 1) * pageSize // Calcula o deslocamento para a paginação

	query := r.connection.Model(model.Filial{}).Preload("Empresa") // Preload adicionado para carregar a empresa

	// Filtra com base no tipo de usuário
	if userType == 2 {
		query = query.Where("criado_por = ?", userID)
	}

	// Conta o total de registros antes da paginação
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("id DESC").Offset(offset).Limit(pageSize).Find(&filial).Error

	if err != nil {
		return nil, 0, err
	}

	return filial, total, nil
}

/** Pega um contrato por ID com validação de tipo de usuário */
func (r *FilialRepository) FindByIDWithUserValidation(id int, userID int, userType int) (*model.Filial, error) {
	var filial model.Filial

	query := r.connection.Model(&model.Filial{}).Preload("Empresa") // Preload adicionado

	// Se o usuário for do tipo 2 (usuário comum), adicione a validação de propriedade
	if userType == 2 {
		query = query.Where("id = ? AND criado_por = ?", id, userID)
	} else {
		// Caso contrário, apenas busque pelo ID (usuário admin)
		query = query.Where("id = ?", id)
	}

	// Executa a consulta
	if err := query.First(&filial).Error; err != nil {
		return nil, err
	}

	return &filial, nil
}
