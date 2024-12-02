package controller

import (
	"net/http"
	"time"

	"go-api/usecase"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(useCase usecase.UserUseCase) *UserController {
	return &UserController{userUseCase: useCase}
}

// Login autentica um usuário e retorna um token
func (ctrl *UserController) Login(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados de login inválidos"})
		return
	}

	user, err := ctrl.userUseCase.Authenticate(loginRequest.Username, loginRequest.Password)
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	// Geração do token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"userID":   user.ID,
		"userType": user.Tipo,                             // Incluindo o userID no payload
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token válido por 24 horas
	})

	tokenString, err := token.SignedString([]byte("secrethash")) // Substitua "secrethash" por uma chave segura
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	// Retorna o token e informações do usuário
	c.JSON(http.StatusOK, gin.H{
		"message": "Login bem-sucedido",
		"token":   tokenString,
		"user": gin.H{
			"name":     user.Name,
			"username": user.Username,
			"type":     user.Tipo,
		},
	})
}

// Register registra um novo usuário
func (ctrl *UserController) Register(c *gin.Context) {
	var registerRequest struct {
		Name     string `json:"name" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Type     int    `json:"type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados de registro inválidos"})
		return
	}

	if registerRequest.Type != 1 && registerRequest.Type != 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de usuário inválido"})
		return
	}

	err := ctrl.userUseCase.RegisterUser(registerRequest.Name, registerRequest.Username, registerRequest.Password, registerRequest.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar usuário",
			"details": err.Error(), // Inclui os detalhes do erro no retorno
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário registrado com sucesso"})
}
