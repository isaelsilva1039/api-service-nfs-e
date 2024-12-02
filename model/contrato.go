package model

import "time"

type Contrato struct {
	ID           int       `gorm:"primaryKey"`
	Nome         string    `gorm:"size:255;not null"`
	CNPJ         string    `gorm:"size:18;not null"`
	CPF          string    `gorm:"size:14;not null"`
	Endereco     string    `gorm:"size:255"`
	PDV          string    `gorm:"size:50"`
	AtivoInativo bool      `gorm:"not null"`
	Telefone     string    `gorm:"size:20"`
	Whatsapp     string    `gorm:"size:20"`
	Email        string    `gorm:"size:255;not null"`
	Responsavel  string    `gorm:"size:255"`
	CriadoPor    int       `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
