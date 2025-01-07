package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ndreyserg/gophermart/internal/api/session"
	"github.com/ndreyserg/gophermart/internal/service"
)

type api struct {
	chi         *chi.Mux
	userService service.UserService
	session     session.Session
}

func NewRouter(userService service.UserService, sessionSecret string) http.Handler {
	a := api{
		chi:         chi.NewRouter(),
		userService: userService,
		session:     session.NewSession(sessionSecret),
	}
	a.chi.Post("/api/user/register", a.Register)
	a.chi.Post("/api/user/login", a.Login)

	return a.chi
}
