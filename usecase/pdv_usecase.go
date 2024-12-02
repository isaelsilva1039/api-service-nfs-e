package usecase

import (
	"errors"
	dto "go-api/dto/pdv"
	"go-api/model"
	"go-api/repository"
)

type PDVUsecase struct {
	pdvRepo *repository.PDVRepository
}

func NewPDVUsecase(pdvRepo *repository.PDVRepository) *PDVUsecase {
	return &PDVUsecase{pdvRepo: pdvRepo}
}

func (u *PDVUsecase) CreatePDV(data dto.CreatePDVRequest, userID int) (*dto.CreatePDVResponse, error) {
	// Validação de negócio, se necessário
	if data.TokenID == nil && (data.DataAtivacaoToken != nil || data.DataExpiracaoToken != nil) {
		return nil, errors.New("não pode definir datas sem um token vinculado")
	}

	pdv := &model.PDV{
		Descricao:          data.Descricao,
		Status:             data.Status,
		FilialID:           data.FilialID,
		ContratoID:         data.ContratoID,
		TokenID:            data.TokenID,
		DataAtivacaoToken:  data.DataAtivacaoToken,
		DataExpiracaoToken: data.DataExpiracaoToken,
		CriadoPor:          userID,
	}

	err := u.pdvRepo.Create(pdv)
	if err != nil {
		return nil, err
	}

	return &dto.CreatePDVResponse{ID: pdv.ID}, nil
}

func (u *PDVUsecase) GetAllPdvs(userID int, userType int, page int, pageSize int) ([]model.PDV, int64, error) {

	return u.pdvRepo.GetAllPdvs(userID, userType, page, pageSize)
}
