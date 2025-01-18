package api

import (
	"errors"
	"net/http"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (a *api) Register(w http.ResponseWriter, r *http.Request) {
	regRequest, err := a.getAuthData(w, r)

	if err != nil {
		return
	}

	user, err := a.userService.Register(r.Context(), regRequest.Login, regRequest.Password)

	if err != nil && errors.Is(err, model.ErrUserAllreadyExist) {
		a.makeErrorResponse(w, "Логин занят", http.StatusConflict)
		return
	}

	if err != nil {
		a.makeErrorResponse(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	a.openSession(user.ID, w, r)
}
