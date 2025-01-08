package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ndreyserg/gophermart/internal/api/session"
	"github.com/ndreyserg/gophermart/internal/service"
)

type api struct {
	chi            *chi.Mux
	userService    service.UserService
	orderService   service.OrderService
	accountService service.AccountService
	session        session.Session
}

func NewRouter(userService service.UserService, orderService service.OrderService, accountService service.AccountService, sessionSecret string) http.Handler {
	a := api{
		chi:            chi.NewRouter(),
		userService:    userService,
		orderService:   orderService,
		accountService: accountService,
		session:        session.NewSession(sessionSecret),
	}
	a.chi.Post("/api/user/register", a.Register)
	a.chi.Post("/api/user/login", a.Login)
	a.chi.Post("/api/user/orders", a.CreateOrder)
	a.chi.Get("/api/user/orders", a.GetUserOrders)
	a.chi.Post("/api/user/balance/withdraw", a.Withdraw)

	return a.chi
}
