package usecase

import (
	"errors"

	dto "go-api/dto/filial"
	"go-api/model"
	"go-api/repository"
)

type FilialUsecase struct {
	filialRepo repository.FilialRepository
}

func NewFilialUsecase(filialRepo repository.FilialRepository) *FilialUsecase {
	return &FilialUsecase{
		filialRepo: filialRepo,
	}
}

/** cria uma empresa */
func (u *FilialUsecase) CreateFilial(data dto.CreateFilialRequest, userID int, userType int) error {

	filial := model.Filial{
		Descricao:         data.Descricao,
		CNPJ:              data.CNPJ,
		InscricaoEstadual: data.InscricaoEstadual,
		RazaoSocial:       data.RazaoSocial,
		NomeFantasia:      data.NomeFantasia,
		Endereco:          data.Endereco,
		FkEmpresa:         data.FkEmpresa,
		CriadoPor:         userID,
		Contribuinte:      data.Contribuinte,
	}

	return u.filialRepo.Save(filial)
}

/** edita uma empresa */
func (u *FilialUsecase) UpdateFilial(id int, data dto.UpdateFilialRequest) error {

	filial, err := u.filialRepo.FindByID(id)
	if err != nil {
		return errors.New("filial não encontrada")
	}

	filial.Descricao = data.Descricao
	filial.CNPJ = data.CNPJ
	filial.InscricaoEstadual = data.InscricaoEstadual
	filial.RazaoSocial = data.RazaoSocial
	filial.NomeFantasia = data.NomeFantasia
	filial.Endereco = data.Endereco
	filial.FkEmpresa = data.FkEmpresa
	filial.Contribuinte = data.Contribuinte

	return u.filialRepo.SaveUpdate(filial)
}

/** remove uma empresa */
func (u *FilialUsecase) RemoveFilial(id int) error {

	filial, err := u.filialRepo.FindByID(id)
	if err != nil {
		return errors.New("Filial não encontrada")
	}

	return u.filialRepo.Delete(filial)
}

/** obtém todas as empresas */
func (u *FilialUsecase) GetFilial(userID int, userType int, page int, pageSize int) ([]model.Filial, int64, error) {
	return u.filialRepo.GetAll(userID, userType, page, pageSize)
}

/** obtém uma empresa pelo ID */
func (u *FilialUsecase) GetFilialByID(id int, userID int, userType int) (*model.Filial, error) {
	return u.filialRepo.FindByIDWithUserValidation(id, userID, userType)
}
