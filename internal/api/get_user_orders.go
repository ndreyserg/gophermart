package api

import (
	"encoding/json"
	"net/http"
)

func (a *api) GetUserOrders(w http.ResponseWriter, r *http.Request) {
	userID, err := a.session.GetUserID(r)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	orders, err := a.orderService.GetByUserID(r.Context(), userID)

	if err != nil {
		a.makeErrorResponse(w, "server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set(contentTypeHeader, contentTypeJSON)

	if len(orders) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	_ = enc.Encode(orders)
}
