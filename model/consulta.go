package model

import "time"

type Consulta struct {
	ID                  int       `json:"id"`
	UserID              int       `json:"user_id"`
	QuantidadeConsultas int       `json:"quantidade_consultas"`
	InicioData          time.Time `json:"inicio_data"`
	FimData             time.Time `json:"fim_data"`
	QuantidadeRealizada int       `json:"quantidade_realizada"`
}

// Define o nome da tabela no GORM
func (Consulta) TableName() string {
	return "consultas"
}
