package model

import "time"

type TokenPdv struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Descricao   string    `gorm:"type:varchar(255);not null" json:"descricao"`
	Status      string    `gorm:"type:varchar(50);not null" json:"status"`
	TokenPDV    string    `gorm:"type:varchar(255);unique;not null" json:"token_pdv"`
	DataCriacao time.Time `gorm:"column:data_criacao;autoCreateTime" json:"data_criacao"`
	CriadoPor   uint      `gorm:"not null" json:"criado_por"`
	User        *User     `gorm:"foreignKey:CriadoPor" json:"user"`
}

// Especifica o nome correto da tabela
func (TokenPdv) TableName() string {
	return "tokens_pdv" // Nome da tabela no banco
}
