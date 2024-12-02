package usecase

import (
	"errors"

	dto "go-api/dto/empresa"
	"go-api/model"
	"go-api/repository"
)

type EmpresaUsecase struct {
	empresaRepo repository.EmpresaRepository
}

func NewEmpresaUsecase(empresaRepo repository.EmpresaRepository) *EmpresaUsecase {
	return &EmpresaUsecase{
		empresaRepo: empresaRepo,
	}
}

/** cria uma empresa */
func (u *EmpresaUsecase) CreateEmpresa(data dto.CreateEmpresaRequest, userID int, userType int) error {

	empresa := model.Empresa{
		Descricao:    data.Descricao,
		RazaoSocial:  data.RazaoSocial,
		NomeFantasia: data.NomeFantasia,
		Endereco:     data.Endereco,
		IDCliente:    userID,
		UserCriacao:  userID,
	}

	return u.empresaRepo.Save(empresa)
}

/** edita uma empresa */
func (u *EmpresaUsecase) UpdateEmpresa(id int, data dto.UpdateEmpresaRequest) error {

	empresa, err := u.empresaRepo.FindByID(id)
	if err != nil {
		return errors.New("empresa não encontrada")
	}

	empresa.Descricao = data.Descricao
	empresa.RazaoSocial = data.RazaoSocial
	empresa.NomeFantasia = data.NomeFantasia
	empresa.Endereco = data.Endereco

	return u.empresaRepo.SaveUpdate(empresa)
}

/** remove uma empresa */
func (u *EmpresaUsecase) RemoveEmpresa(id int) error {

	empresa, err := u.empresaRepo.FindByID(id)
	if err != nil {
		return errors.New("empresa não encontrada")
	}

	return u.empresaRepo.Delete(empresa)
}

/** obtém todas as empresas */
func (u *EmpresaUsecase) GetEmpresas(userID int, userType int, page int, pageSize int) ([]model.Empresa, int64, error) {
	return u.empresaRepo.GetAll(userID, userType, page, pageSize)
}

/** obtém uma empresa pelo ID */
func (u *EmpresaUsecase) GetEmpresaByID(id int, userID int, userType int) (*model.Empresa, error) {
	return u.empresaRepo.FindByIDWithUserValidation(id, userID, userType)
}
