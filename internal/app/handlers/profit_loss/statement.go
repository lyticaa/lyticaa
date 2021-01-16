package profit_loss

import (
	"encoding/json"
	"net/http"

	"github.com/lyticaa/lyticaa/internal/app/helpers"
	"github.com/lyticaa/lyticaa/internal/app/types"
)

func (p *ProfitLoss) Statement(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(p.sessionStore, p.logger, w, r)
	helpers.RenderTemplate(w, helpers.AppLayout, helpers.TemplateList(helpers.ProfitLossStatement), session.Values)
}

func (p *ProfitLoss) StatementByDate(w http.ResponseWriter, r *http.Request) {
	_ = helpers.GetSessionUser(helpers.GetSession(p.sessionStore, p.logger, w, r))

	var byDate types.Statement
	byDate.Data = []types.StatementTable{}
	byDate.Draw = helpers.DtDraw(r)

	js, err := json.Marshal(byDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
