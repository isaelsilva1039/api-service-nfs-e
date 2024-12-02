package model

import "time"

type PDV struct {
	ID                 uint       `gorm:"primaryKey" json:"id"`
	Descricao          string     `gorm:"type:varchar(255);not null" json:"descricao"`
	Status             string     `gorm:"type:varchar(50);not null" json:"status"`
	DataCriacao        time.Time  `gorm:"column:data_criacao;autoCreateTime" json:"data_criacao"`
	DataAtivacaoToken  *time.Time `gorm:"column:data_ativacao_token" json:"data_ativacao_token"`
	DataExpiracaoToken *time.Time `gorm:"column:data_expiracao_token" json:"data_expiracao_token"`
	CriadoPor          int        `gorm:"column:criado_por;not null" json:"criado_por"`
	User               *User      `gorm:"foreignKey:CriadoPor" json:"user"`
	FilialID           uint       `gorm:"not null" json:"filial_id"`
	Filial             *Filial    `gorm:"foreignKey:FilialID" json:"filial"`
	ContratoID         uint       `gorm:"not null" json:"contrato_id"`
	Contrato           *Contrato  `gorm:"foreignKey:ContratoID" json:"contrato"`
	TokenID            *uint      `gorm:"column:token_id" json:"token_id"`
	Token              *TokenPdv  `gorm:"foreignKey:TokenID" json:"token"`
}
