package api

import (
	"net/http"
)

func (a *api) openSession(userID int, w http.ResponseWriter, r *http.Request) {
	err := a.session.Open(userID, w, r)

	if err != nil {
		a.makeErrorResponse(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
