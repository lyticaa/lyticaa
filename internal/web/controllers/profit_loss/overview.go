package profit_loss

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/helpers"
	"gitlab.com/getlytica/lytica-app/internal/web/types"
)

func (p *ProfitLoss) Overview(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(p.sessionStore, p.logger, w, r)

	t := []string{
		"partials/_nav",
		"partials/nav/_main",
		"partials/nav/account/_main",
		"partials/filters/_filters",
		"partials/filters/_date",
		"partials/filters/_upload",
		"profit_loss/overview",
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
