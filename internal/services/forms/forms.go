package forms

import (
	"FormManager/internal/model"
	"FormManager/internal/services/authorisation"
	"FormManager/internal/services/jwt"
	"fmt"
	"math/rand"
)

type FormRepository interface {
	CreateForm(form model.Form, formID uint) (uint, error)
	EditForm(form model.Form) error
	DeleteForm(formID uint) error
	GetFormByID(formID uint) (*model.Form, error)
}

type FormService struct {
	FormRepository FormRepository
	UserRepository authorisation.UserRepository //Спросить про это
}

func NewFormService(formRepo FormRepository, userRepo authorisation.UserRepository) *FormService {
	return &FormService{FormRepository: formRepo, UserRepository: userRepo}
}

func (f *FormService) CreateForm(form model.Form, token string) (uint, error) {
	jwtClaims, err := jwt.ParseToken(token)
	if err != nil {
		return 0, fmt.Errorf("failed to parse token: %w", err)
	}
	user, err := f.UserRepository.GetUserInformation(jwtClaims.Email)
	if err != nil {
		return 0, fmt.Errorf("failed to get user information: %w", err)
	}
	form.Creator = *user
	uid, err := f.FormRepository.CreateForm(form, generateRandomID())
	if err != nil {
		return 0, err
	}
	return uid, nil
}

func (f *FormService) EditForm(form model.Form, token string) error {
	jwtClaims, err := jwt.ParseToken(token)
	if err != nil {
		return fmt.Errorf("failed to parse token: %w", err)
	}

	user, err := f.UserRepository.GetUserInformation(jwtClaims.Email)
	if err != nil {
		return fmt.Errorf("failed to get user information: %w", err)
	}

	if user.Role.HasPermission(model.PermissionDelete) || user.GetEmail() == form.Creator.GetEmail() {
		return f.FormRepository.EditForm(form)
	}
	return fmt.Errorf("unauthorized to edit form")
}

func (f *FormService) DeleteForm(formID uint, token string) error {
	jwtClaims, err := jwt.ParseToken(token)
	if err != nil {
		return fmt.Errorf("failed to parse token: %w", err)
	}

	user, err := f.UserRepository.GetUserInformation(jwtClaims.Email)
	if err != nil {
		return fmt.Errorf("failed to get user information: %w", err)
	}

	form, err := f.FormRepository.GetFormByID(formID)
	if err != nil {
		return err
	}

	if user.Role.HasPermission(model.PermissionDelete) || user.GetEmail() == form.Creator.GetEmail() {
		return f.FormRepository.DeleteForm(formID)
	}
	return fmt.Errorf("unauthorized to delete form")
}

func (f *FormService) GetForm(formID uint) (*model.Form, error) {
	form, err := f.FormRepository.GetFormByID(formID)
	if err != nil {
		return &model.Form{}, err
	}
	return form, nil
}
func generateRandomID() uint {
	return uint(rand.Uint32())
}
