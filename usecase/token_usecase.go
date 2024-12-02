package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type TokenUsecase struct {
	tokenRepo *repository.TokenRepository
}

func NewTokenUsecase(tokenRepo *repository.TokenRepository) *TokenUsecase {
	return &TokenUsecase{tokenRepo: tokenRepo}
}

func (u *TokenUsecase) GetAllTokens(userID int, userType int, page int, pageSize int) ([]model.TokenPdv, int64, error) {

	return u.tokenRepo.GetAllTokens(userID, userType, page, pageSize)
}
