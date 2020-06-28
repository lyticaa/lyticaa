package api

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/web/app/types"
)

type API struct{}

func NewAPI() *API {
	return &API{}
}

func (a *API) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := types.Health{Status: "OK"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
