package admin

import (
	"context"

	"github.com/lyticaa/lyticaa-app/internal/app/models"
	"github.com/lyticaa/lyticaa-app/internal/app/types"

	"github.com/jmoiron/sqlx"
)

func Users(ctx context.Context, admin *types.Admin, filter *models.Filter, db *sqlx.DB) {
	var userModel models.UserModel
	users := userModel.FetchAll(ctx, nil, filter, db).([]models.UserModel)

	for _, user := range users {
		table := types.AdminTable{
			RowID:   user.UserID,
			Email:   user.Email,
			Created: user.CreatedAt.Format("2006-01-02"),
		}

		admin.Data = append(admin.Data, table)
	}

	admin.RecordsTotal = userModel.Count(ctx, nil, db)
	admin.RecordsFiltered = admin.RecordsTotal
}
