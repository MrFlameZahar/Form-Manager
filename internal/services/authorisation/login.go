package authorisation

import (
	"FormManager/internal/services/jwt"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (a Auth) LoginUser(email string, password string) (string, error) {
	user, err := a.UserRepository.GetUserInformation(email)

	if err != nil {
		return "", err
	}

	if user == nil {
		return "", fmt.Errorf("user not found")
	}
	passwordHash, err := a.UserRepository.GetPasswordHash(uint(user.ID))
	if err != nil {
		return "", fmt.Errorf("error: %w", err)
	}
	if err := bcrypt.CompareHashAndPassword(passwordHash, []byte(password)); err != nil {
		return "", fmt.Errorf("wrong login or password")
	}

	token, err := jwt.GenerateJWT(user.GetEmail())
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %w", err)
	}
	return token, nil
}
