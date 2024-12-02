package model

import "time"

// Empresa representa uma entidade empresarial
type Empresa struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Descricao    string    `gorm:"type:text" json:"descricao"`
	RazaoSocial  string    `gorm:"type:varchar(255)" json:"razao_social"`
	NomeFantasia string    `gorm:"type:varchar(255)" json:"nome_fantasia"`
	Endereco     string    `gorm:"type:varchar(255)" json:"endereco"`
	IDCliente    int       `gorm:"column:id_cliente" json:"id_cliente"`
	DataCriacao  time.Time `gorm:"column:criado_em;autoCreateTime" json:"criado_em"`
	UserCriacao  int       `gorm:"column:fk_user_criacao" json:"fk_user_criacao"`
}
