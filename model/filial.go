package model

import "time"

// Filial representa uma entidade filial vinculada a uma empresa
type Filial struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	Descricao         string    `gorm:"type:varchar(255)" json:"descricao"`
	CNPJ              string    `gorm:"type:varchar(18);uniqueIndex" json:"cnpj"` // Formato CNPJ com máscara
	InscricaoEstadual string    `gorm:"type:varchar(20)" json:"inscricao_estadual"`
	RazaoSocial       string    `gorm:"type:varchar(255)" json:"razao_social"`
	NomeFantasia      string    `gorm:"type:varchar(255)" json:"nome_fantasia"`
	Endereco          string    `gorm:"type:varchar(255)" json:"endereco"`
	FkEmpresa         uint      `gorm:"column:fk_empresa;not null" json:"fk_empresa"` // ID da empresa associada
	Empresa           Empresa   `gorm:"foreignKey:FkEmpresa;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"empresa"`
	CriadoEm          time.Time `gorm:"column:criado_em;autoCreateTime" json:"criado_em"`
	CriadoPor         int       `gorm:"not null" json:"criado_por"`        // Representa o ID do usuário logado
	Contribuinte      bool      `gorm:"default:false" json:"contribuinte"` // Indica se é contribuinte ou não

}
