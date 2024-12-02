package usecase

import (
	"errors"
	dto "go-api/dto/contrato"
	"go-api/model"
	"go-api/repository"
)

type ContratoUsecase struct {
	contratoRepo repository.ContratoRepository
}

func NewContratoUsecase(contratoRepo repository.ContratoRepository) *ContratoUsecase {
	return &ContratoUsecase{
		contratoRepo: contratoRepo,
	}

}

/** cria um contrato */
func (u *ContratoUsecase) CreateContrato(data dto.CreateContratoRequest, userID int) error {
	// Converte os dados recebidos no DTO para o modelo da base de dados
	contrato := model.Contrato{
		Nome:         data.Nome,
		CNPJ:         data.CNPJ,
		CPF:          data.CPF,
		Endereco:     data.Endereco,
		PDV:          data.PDV,
		AtivoInativo: data.AtivoInativo,
		Telefone:     data.Telefone,
		Whatsapp:     data.Whatsapp,
		Email:        data.Email,
		Responsavel:  data.Responsavel,
		CriadoPor:    userID, // Adiciona o ID do usuário logado
	}

	// Salva no repositório
	return u.contratoRepo.Save(contrato)
}

/** edita um contrato */
func (u *ContratoUsecase) UpdateContrato(id int, data dto.UpdateContratoRequest, userID int) error {
	// Busque o contrato para verificar a existência
	contrato, err := u.contratoRepo.FindByID(id)
	if err != nil {
		return errors.New("contrato não encontrado")
	}

	// Verifique se o usuário logado é o criador do contrato
	if contrato.CriadoPor != userID {
		return errors.New("usuário não autorizado a editar este contrato")
	}

	// Atualize os campos permitidos
	contrato.Nome = data.Nome
	contrato.CNPJ = data.CNPJ
	contrato.CPF = data.CPF
	contrato.Endereco = data.Endereco
	contrato.PDV = data.PDV
	contrato.AtivoInativo = data.AtivoInativo
	contrato.Telefone = data.Telefone
	contrato.Whatsapp = data.WhatsApp
	contrato.Email = data.Email
	contrato.Responsavel = data.Responsavel

	// Salve o contrato atualizado
	return u.contratoRepo.SaveUpdate(contrato)
}

/** remove um contatos */
func (u *ContratoUsecase) Remove(id int) error {

	contrato, err := u.contratoRepo.FindByID(id)
	if err != nil {
		return errors.New("contrato não encontrado")
	}

	// Salve o contrato atualizado
	return u.contratoRepo.Delete(contrato)
}

/** obtem os contratos */
func (u *ContratoUsecase) GetContratos(userID int, userType int, page int, pageSize int) ([]model.Contrato, int64, error) {

	return u.contratoRepo.GetAll(userID, userType, page, pageSize)

}

/** obtem os contratos */
func (u *ContratoUsecase) GetContratoByID(id int, userID int, userType int) (*model.Contrato, error) {

	return u.contratoRepo.FindByIDWithUserValidation(id, userID, userType)

}
