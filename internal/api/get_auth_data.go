package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

type req struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

var errBadRequest = errors.New("bad request")

func (a *api) getAuthData(w http.ResponseWriter, r *http.Request) (*req, error) {
	req := req{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		a.makeErrorResponse(w, "Неверный запрос", http.StatusBadRequest)
		return nil, errBadRequest
	}

	if req.Login == "" {
		a.makeErrorResponse(w, "Не указан логин", http.StatusBadRequest)
		return nil, errBadRequest
	}

	if req.Password == "" {
		a.makeErrorResponse(w, "Не указан пароль", http.StatusBadRequest)
		return nil, errBadRequest
	}
	return &req, nil
}
