package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	apperror "github.com/nico-hondo/internal/error"
	"github.com/nico-hondo/internal/model"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (ur *UserRepo) GetProfileUserByID(ctx context.Context, id int) (model.Account, error) {
	sql := "SELECT id, name, email, avatar_url, bio, location, role, created_at from users WHERE id=$1"
	args := []any{id}

	var data model.Account
	if err := ur.db.QueryRow(ctx, sql, args...).Scan(&data.Id, &data.Name, &data.Email, &data.AvatarURL, &data.Bio, &data.Location, &data.Role, &data.CreatedAt); err != nil {
		return model.Account{}, err
	}

	return data, nil
}

func (ur *UserRepo) ChangeUserProfile(ctx context.Context, name, avatarURL, bio, location string, id int) error {
	sql := `
		UPDATE users
		SET 
			name = COALESCE(NULLIF($1, ''), name),
			avatar_url = COALESCE(NULLIF($2, ''), avatar_url),
			bio = COALESCE(NULLIF($3, ''), bio),
			location = COALESCE(NULLIF($4, ''), location),
			updated_at = NOW()
		WHERE id = $5`

	args := []any{name, avatarURL, bio, location, id}

	cmd, err := ur.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return apperror.ErrNotFound
	}

	return nil
}
