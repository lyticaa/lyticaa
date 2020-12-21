package users

import (
	"context"
	"database/sql"

	"github.com/lyticaa/lyticaa-app/internal/models"

	"github.com/jmoiron/sqlx"
)

func User(ctx context.Context, userID string, db *sqlx.DB) models.UserModel {
	userModel := models.UserModel{
		UserID: userID,
	}

	user := userModel.FetchOne(ctx, db).(models.UserModel)

	return user
}

func Create(ctx context.Context, userID, email, nickname, avatarURL string, db *sqlx.DB) error {
	user := models.UserModel{
		UserID: userID,
		Email:  email,
	}

	var userNickname sql.NullString
	if err := userNickname.Scan(nickname); err != nil {
		return err
	}
	user.Nickname = userNickname

	var userAvatarURL sql.NullString
	if err := userAvatarURL.Scan(avatarURL); err != nil {
		return err
	}
	user.AvatarURL = userAvatarURL

	if err := user.Create(ctx, db); err != nil {
		return err
	}

	return nil
}
