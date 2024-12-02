package usecase

import (
	"go-api/model"
	"go-api/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Authenticate(username, password string) (*model.User, error)
	RegisterUser(name string, username string, password string, userType int) error
}

type userUseCase struct {
	userRepo repository.UserRepository
}

// NewUserUseCase cria uma nova instância do caso de uso de usuários
func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

// Authenticate autentica o usuário e retorna os dados completos do usuário se válido
func (u *userUseCase) Authenticate(username, password string) (*model.User, error) {
	user, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	// Verifica a senha
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, nil
	}

	// Retorna o usuário se autenticado
	return user, nil
}

// RegisterUser registra um novo usuário com senha criptografada
func (u *userUseCase) RegisterUser(name string, username string, password string, userType int) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Cria um novo usuário
	user := &model.User{
		Username: username,
		Password: string(hashedPassword),
		Name:     name,
		Tipo:     userType,
	}

	// Salva no repositório
	return u.userRepo.CreateUser(user)
}
