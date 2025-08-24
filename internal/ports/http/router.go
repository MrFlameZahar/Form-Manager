package router

import (
	authHandler "FormManager/internal/ports/http/auth"
	"FormManager/internal/services/authorisation"

	"github.com/go-chi/chi"
)

func InitializeRouter(authService authorisation.Auth) *chi.Mux {
	router := chi.NewRouter()

	router.Post("/auth/login", authHandler.Login(authService))
	router.Post("/auth/signup", authHandler.Signup(authService))
	router.Get("/auth/me", authHandler.Me(authService))

	router.Post("/forms/create", nil)
	router.Put("/forms/{id}/edit", nil)
	router.Delete("/forms/{id}/delete", nil)
	router.Get("/forms/{id}", nil)

	router.Get("/forms/{id}/responses", nil)
	router.Get("/forms/{id}/responses/{id}", nil)
	router.Post("/forms/{id}/responses", nil)

	router.Get("/user/{id}/forms", nil)
	router.Get("/user", nil)
	return router
}
