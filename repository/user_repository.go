package repository

import (
	"go-api/model"

	"gorm.io/gorm"
)

// UserRepository define a interface para o repositório de usuários
type UserRepository interface {
	FindByUsername(username string) (*model.User, error)
	CreateUser(user *model.User) error
}

// userRepository implementa a interface UserRepository
type userRepository struct {
	connection *gorm.DB
}

// NewUserRepository cria uma nova instância de userRepository
func NewUserRepository(connection *gorm.DB) UserRepository {
	return &userRepository{connection: connection}
}

func (r *userRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.connection.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user *model.User) error {
	return r.connection.Create(user).Error
}
