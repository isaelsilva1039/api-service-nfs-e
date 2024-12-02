package dto

// CreateFilialRequest representa a requisição para criar uma nova filial
type CreateFilialRequest struct {
	Descricao         string `json:"descricao" binding:"required"`
	CNPJ              string `json:"cnpj" binding:"required"`
	InscricaoEstadual string `json:"inscricao_estadual"`
	RazaoSocial       string `json:"razao_social" binding:"required"`
	NomeFantasia      string `json:"nome_fantasia" binding:"required"`
	Endereco          string `json:"endereco" binding:"required"`
	FkEmpresa         uint   `json:"fk_empresa" binding:"required"`
	// CriadoPor         int    `json:"criado_por"`
	Contribuinte bool `json:"contribuinte"`
}

// UpdateFilialRequest representa a requisição para atualizar uma filial existente
type UpdateFilialRequest struct {
	Descricao         string `json:"descricao" binding:"required"`
	CNPJ              string `json:"cnpj" binding:"required"`
	InscricaoEstadual string `json:"inscricao_estadual"`
	RazaoSocial       string `json:"razao_social" binding:"required"`
	NomeFantasia      string `json:"nome_fantasia" binding:"required"`
	Endereco          string `json:"endereco" binding:"required"`
	FkEmpresa         uint   `json:"fk_empresa" binding:"required"`
	// CriadoPor         int    `json:"criado_por"`
	Contribuinte bool `json:"contribuinte"`
}
