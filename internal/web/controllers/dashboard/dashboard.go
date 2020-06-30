package dashboard

import (
	"gitlab.com/getlytica/lytica-app/internal/web/lib/data"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Dashboard struct {
	data         *data.Data
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewDashboard(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *Dashboard {
	return &Dashboard{
		data:         data.NewData(db),
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
	}
}
