package api

import (
	"encoding/json"
	"net/http"
)

func (a *api) GetUserWithdrawals(w http.ResponseWriter, r *http.Request) {
	userID, err := a.session.GetUserID(r)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	withdrawals, err := a.accountService.GetWithdrawals(r.Context(), userID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(withdrawals) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set(contentTypeHeader, contentTypeJSON)
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	_ = e.Encode(withdrawals)
}
