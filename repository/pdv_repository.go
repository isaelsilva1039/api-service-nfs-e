package repository

import (
	"go-api/model"

	"gorm.io/gorm"
)

type PDVRepository struct {
	db *gorm.DB
}

func NewPDVRepository(db *gorm.DB) *PDVRepository {
	return &PDVRepository{db: db}
}

func (r *PDVRepository) Create(pdv *model.PDV) error {
	return r.db.Create(pdv).Error
}

/*
  - busca os PDV
    Tipo: 1 pega todos os PDV (perfil adm)
    Tipo: 2 pega só os PDV que ele cadastrou (perfil parceiro)
*/
func (r *PDVRepository) GetAllPdvs(userID int, userType int, page int, pageSize int) ([]model.PDV, int64, error) {
	var pdv []model.PDV
	var total int64

	offset := (page - 1) * pageSize // Calcula o deslocamento para a paginação

	query := r.db.Model(model.PDV{}).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "tipo") // Seleciona apenas os campos necessários do usuário
		}).
		Preload("Filial", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Empresa") // Preload da empresa ao carregar a filial
		}).
		Preload("Contrato"). // Carrega os dados do contrato associado
		Preload("Token")     // Carrega os dados do token associado

	// Filtra com base no tipo de usuário
	if userType == 2 {
		query = query.Where("criado_por = ?", userID)
	}

	// Conta o total de registros antes da paginação
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("id DESC").Offset(offset).Limit(pageSize).Find(&pdv).Error

	if err != nil {
		return nil, 0, err
	}

	return pdv, total, nil
}
