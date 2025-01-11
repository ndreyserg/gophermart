package api

import (
	"encoding/json"
	"net/http"
)

func (a *api) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	userID, err := a.session.GetUserID(r)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	balance, err := a.accountService.GetBalance(r.Context(), userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set(contentTypeHeader, contentTypeJSON)
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	_ = e.Encode(balance)
}
