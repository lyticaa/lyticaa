package profit_loss

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/models"
)

func (p *ProfitLoss) Overview(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(p.sessionStore, p.logger, w, r)

	t := []string{
		"partials/nav/_main",
		"profit_loss/overview",
		"partials/_filters",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (p *ProfitLoss) ProfitLossByDate(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(p.sessionStore, p.logger, w, r)
	_ = session.Values["User"].(models.User)

	table := []types.ProfitLossTable{}
	byDate := types.ProfitLoss{Data: table}

	js, err := json.Marshal(byDate)
	if err != nil {
		p.logger.Error().Err(err).Msg("unable to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
