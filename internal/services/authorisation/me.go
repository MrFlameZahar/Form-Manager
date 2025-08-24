package authorisation

import (
	"FormManager/internal/model"
	"FormManager/internal/services/jwt"
	"fmt"
)

func (a Auth) Me(token string) (*model.User, error) {
	jwtClaims, err := jwt.ParseToken(token)
	if err != nil {
		return &model.User{}, fmt.Errorf("failed to parse token: %w", err)
	}
	user, err := a.UserRepository.GetUserInformation(jwtClaims.Email)
	if err != nil {
		return &model.User{}, fmt.Errorf("failed to get user information: %w", err)
	}
	return user, nil
}
