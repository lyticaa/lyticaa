package profit_loss

import (
	"encoding/json"
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/models"
	"github.com/lyticaa/lyticaa-app/internal/web/helpers"
	"github.com/lyticaa/lyticaa-app/internal/web/types"
)

func (p *ProfitLoss) Overview(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(p.sessionStore, p.logger, w, r)
	helpers.RenderTemplate(w, helpers.TemplateList(helpers.ProfitLossOverview), session.Values)
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
