package dto

import "time"

// CreatePDVRequest representa os dados para criar um PDV
type CreatePDVRequest struct {
	Descricao          string     `json:"descricao" validate:"required"`
	Status             string     `json:"status" validate:"required"`
	FilialID           uint       `json:"filial_id" validate:"required"`
	ContratoID         uint       `json:"contrato_id" validate:"required"`
	TokenID            *uint      `json:"token_id"` // Opcional
	DataAtivacaoToken  *time.Time `json:"data_ativacao_token"`
	DataExpiracaoToken *time.Time `json:"data_expiracao_token"`
}

// CreatePDVResponse representa a resposta após criar um PDV
type CreatePDVResponse struct {
	ID uint `json:"id"`
}
