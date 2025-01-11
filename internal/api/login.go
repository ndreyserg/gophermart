package api

import (
	"errors"
	"net/http"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (a *api) Login(w http.ResponseWriter, r *http.Request) {
	logReq, err := a.getAuthData(w, r)

	if err != nil {
		return
	}

	user, err := a.userService.Login(r.Context(), logReq.Login, logReq.Password)

	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			a.makeErrorResponse(w, "Неверная пара логин/пароль", http.StatusUnauthorized)
			return
		}
		a.makeErrorResponse(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}
	a.openSession(user.ID, w, r)
}
