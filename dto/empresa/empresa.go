package dto

// CreateEmpresaRequest representa a requisição para criar uma nova empresa
type CreateEmpresaRequest struct {
	Descricao    string `json:"descricao" binding:"required"`
	RazaoSocial  string `json:"razao_social" binding:"required"`
	NomeFantasia string `json:"nome_fantasia" binding:"required"`
	Endereco     string `json:"endereco" binding:"required"`
	// UserCriacao  string `json:"fk_user_criacao"`
}

// UpdateEmpresaRequest representa a requisição para atualizar uma empresa existente
type UpdateEmpresaRequest struct {
	Descricao    string `json:"descricao" binding:"required"`
	RazaoSocial  string `json:"razao_social" binding:"required"`
	NomeFantasia string `json:"nome_fantasia" binding:"required"`
	Endereco     string `json:"endereco" binding:"required"`
	// UserCriacao  string `json:"fk_user_criacao"`
}
