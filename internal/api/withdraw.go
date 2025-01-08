package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ndreyserg/gophermart/internal/model"
)

type withdrawRequest struct {
	Order string
	Sum   float64
}

func (a *api) Withdraw(w http.ResponseWriter, r *http.Request) {

	userID, err := a.session.GetUserID(r)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	req := withdrawRequest{}
	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&req)
	if err != nil {
		a.makeErrorResponse(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	if req.Order == "" {
		a.makeErrorResponse(w, "Неверный номер заказа", http.StatusUnprocessableEntity)
		return
	}

	if req.Sum <= 0 {
		a.makeErrorResponse(w, "Неверная сумма", http.StatusUnprocessableEntity)
		return
	}

	err = a.accountService.Withdraw(r.Context(), userID, req.Order, req.Sum)

	if err != nil {
		if errors.Is(err, model.ErrorAccountNegativeBalance) {
			w.WriteHeader(http.StatusPaymentRequired)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
