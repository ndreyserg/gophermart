package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ndreyserg/gophermart/internal/model"
)

type loginRequest struct {
	Login    string
	Password string
}

func (a *api) Login(w http.ResponseWriter, r *http.Request) {
	req := loginRequest{}

	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&req); err != nil {
		a.makeErrorResponse(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	if req.Login == "" {
		a.makeErrorResponse(w, "Не указан логин", http.StatusBadRequest)
		return
	}

	if req.Password == "" {
		a.makeErrorResponse(w, "Не указан пароль", http.StatusBadRequest)
		return
	}

	user, err := a.userService.Login(r.Context(), req.Login, req.Password)

	if err != nil {
		if errors.Is(err, model.ErrorUserNotFound) {
			a.makeErrorResponse(w, "Неверная пара логин/пароль", http.StatusUnauthorized)
			return
		}
		a.makeErrorResponse(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	a.session.Open(user.ID, w, r)
	w.WriteHeader(http.StatusOK)
}
