package expenses

import (
	"gitlab.com/lyticaa/lyticaa-app/internal/web/lib/data"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type ValidateExpense struct {
	ExpenseId string `validate:"required,uuid4"`
}

type Expenses struct {
	data         *data.Data
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewExpenses(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *Expenses {
	return &Expenses{
		data:         data.NewData(db),
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
	}
}
