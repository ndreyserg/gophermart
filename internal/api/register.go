package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ndreyserg/gophermart/internal/model"
)

type registerRequest struct {
	Login    string
	Password string
}

func (a *api) Register(w http.ResponseWriter, r *http.Request) {

	req := registerRequest{}
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
		a.makeErrorResponse(w, "Не указан пароль ", http.StatusBadRequest)
		return
	}

	user, err := a.userService.Register(r.Context(), req.Login, req.Password)

	if err != nil {
		if errors.Is(err, model.ErrorUserAllreadyExist) {
			a.makeErrorResponse(w, "Логин занят", http.StatusConflict)
			return
		}
		a.makeErrorResponse(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	err = a.session.Open(user.ID, w, r)

	if err != nil {
		a.makeErrorResponse(w, "Ошибка сервера", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
