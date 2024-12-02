package model

type Cliente struct {
	ID                   int        `json:"id"`
	Name                 string     `json:"name"`
	Email                string     `json:"email"`
	Phone                string     `json:"phone"`
	MobilePhone          string     `json:"mobile_phone"`
	CpfCnpj              string     `json:"cpfCnpj" gorm:"column:cpfCnpj"`
	PostalCode           string     `json:"postal_code"`
	Address              string     `json:"address"`
	AddressNumber        string     `json:"address_number"`
	Complement           string     `json:"complement"`
	Province             string     `json:"province"`
	ExternalReference    string     `json:"external_reference"`
	NotificationDisabled bool       `json:"notification_disabled"`
	Observations         string     `json:"observations"`
	IDClienteAsaas       string     `json:"id_cliente_asaas"`
	DateOfBirth          string     `json:"date_of_birth"`
	UserID               int        `json:"user_id"`
	Total                float64    `json:"total"`
	UpdatedAt            string     `json:"updated_at"`
	CreatedAt            string     `json:"created_at"`
	FKPlano              int        `json:"fk_plano"`
	Inadimplente         bool       `json:"inadimplente"`
	Consultas            []Consulta `json:"consultas" gorm:"foreignKey:UserID;references:UserID"` // Define a relação
}

// PaginationMeta estrutura para metadados de paginação
type PaginationMeta struct {
	MaxPerPage int `json:"max_per_page"`
	Page       int `json:"page"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

// Response estrutura para encapsular os dados e metadados de paginação
type ResponseCliente struct {
	Data []Cliente      `json:"data"`
	Meta PaginationMeta `json:"meta"`
}
