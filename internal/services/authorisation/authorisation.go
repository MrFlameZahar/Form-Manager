package authorisation

import "FormManager/internal/model"

type Auth struct {
	SignupRepository SignupRepository
	UserRepository   UserRepository
}

type UserRepository interface {
	GetUserInformation(email string) (*model.User, error)
	GetPasswordHash(userID uint) ([]byte, error)
}

func NewAuth(signup SignupRepository, login UserRepository) *Auth {
	return &Auth{SignupRepository: signup, UserRepository: login}
}
