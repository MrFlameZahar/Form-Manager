package formsHandler

import (
	"FormManager/internal/model"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

type FormInterface interface {
	CreateForm(form model.Form, token string) (uint, error)
	EditForm(form model.Form, token string) error
	DeleteForm(formID uint, token string) error
	GetForm(formID uint) (model.Form, error)
}

func CreateForm(form FormInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")

		var userForm model.Form

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&userForm); err != nil {
			http.Error(w, "Failed to parse JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		formID, err := form.CreateForm(userForm, token)
		if err != nil {
			http.Error(w, "Create faild: "+err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, formID)
	}
}

func EditForm(form FormInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation goes here
	}
}

func DeleteForm(form FormInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation goes here
	}
}
func GetForm(form FormInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation goes here
	}
}
