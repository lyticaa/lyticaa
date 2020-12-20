package users

import (
	"context"
	"database/sql"

	"github.com/lyticaa/lyticaa-app/internal/models"

	"github.com/jmoiron/sqlx"
)

func FetchUser(ctx context.Context, userID string, db *sqlx.DB) models.UserModel {
	var userModel *models.UserModel

	userModel.UserID = userID
	user := (*userModel).FetchOne(ctx, db)

	return user.(models.UserModel)
}

func CreateUser(ctx context.Context, userID, email, nickname, avatarURL string, db *sqlx.DB) error {
	userModel := &models.UserModel{
		UserID: userID,
		Email:  email,
	}

	var userNickname sql.NullString
	if err := userNickname.Scan(nickname); err != nil {
		return err
	}
	userModel.Nickname = userNickname

	var userAvatarURL sql.NullString
	if err := userAvatarURL.Scan(avatarURL); err != nil {
		return err
	}
	userModel.AvatarURL = userAvatarURL

	if err := (*userModel).Create(ctx, db); err != nil {
		return err
	}

	return nil
}
