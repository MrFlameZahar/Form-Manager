package authorisation

import (
	"golang.org/x/crypto/bcrypt"
)

type SignupRepository interface {
	SaveUserInformation(email string, username string, passwordHash []byte) error
}

func (a Auth) SignupUser(username string, email string, password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = a.SignupRepository.SaveUserInformation(email, username, passwordHash)
	if err != nil {
		return err
	}
	return nil
}
