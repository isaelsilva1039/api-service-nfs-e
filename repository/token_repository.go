package repository

import (
	"go-api/model"

	"gorm.io/gorm"
)

type TokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

/*
  - busca os PDV
    Tipo: 1 pega todos os PDV (perfil adm)
    Tipo: 2 pega só os PDV que ele cadastrou (perfil parceiro)
*/
func (r *TokenRepository) GetAllTokens(userID int, userType int, page int, pageSize int) ([]model.TokenPdv, int64, error) {
	var token []model.TokenPdv
	var total int64

	offset := (page - 1) * pageSize // Calcula o deslocamento para a paginação
	query := r.db.Model(&model.TokenPdv{}).
		Select("id, descricao, status, data_criacao, criado_por").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, tipo")
		})

	// Filtra com base no tipo de usuário
	if userType == 2 {
		query = query.Where("criado_por = ?", userID)
	}

	// Conta o total de registros antes da paginação
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("id DESC").Offset(offset).Limit(pageSize).Find(&token).Error

	if err != nil {
		return nil, 0, err
	}

	return token, total, nil
}
