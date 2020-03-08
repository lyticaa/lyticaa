package app

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica/internal/core/app/types"
)

func (a *App) healthCheck(w http.ResponseWriter, r *http.Request) {
	response := types.Health{Status: "OK"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		a.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonResponse)
	if err != nil {
		a.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
