package api

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (a *api) CreateOrder(w http.ResponseWriter, r *http.Request) {
	userID, err := a.session.GetUserID(r)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	b, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = a.orderService.Create(r.Context(), strings.Trim(string(b), " "), userID)

	if err != nil {
		if errors.Is(err, model.ErrUncorrectOrederNumber) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		if errors.Is(err, model.ErrOrderAllreadyExistOnUser) {
			w.WriteHeader(http.StatusOK)
			return
		}

		if errors.Is(err, model.ErrOrderAllreadyExist) {
			w.WriteHeader(http.StatusConflict)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
