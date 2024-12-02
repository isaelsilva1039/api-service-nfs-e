package dto

type CreateContratoRequest struct {
	Nome         string `json:"nome" binding:"required"`
	CNPJ         string `json:"cnpj" binding:"required"`
	CPF          string `json:"cpf" binding:"required"`
	Endereco     string `json:"endereco" binding:"required"`
	PDV          string `json:"pdv" binding:"required"`
	AtivoInativo bool   `json:"ativo_inativo" binding:"required"`
	Telefone     string `json:"telefone" binding:"required"`
	Whatsapp     string `json:"whatsapp"`
	Email        string `json:"email" binding:"required,email"`
	Responsavel  string `json:"responsavel" binding:"required"`
}

type UpdateContratoRequest struct {
	Nome         string `json:"nome" binding:"required"`
	CNPJ         string `json:"cnpj" binding:"required"`
	CPF          string `json:"cpf" binding:"required"`
	Endereco     string `json:"endereco" binding:"required"`
	PDV          string `json:"pdv" binding:"required"`
	AtivoInativo bool   `json:"ativo_inativo"`
	Telefone     string `json:"telefone" binding:"required"`
	WhatsApp     string `json:"whatsapp"`
	Email        string `json:"email" binding:"required,email"`
	Responsavel  string `json:"responsavel" binding:"required"`
}
