package repo

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	apperror "github.com/nico-hondo/internal/error"
	"github.com/nico-hondo/internal/model"
)

type AuthRepo struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (ar *AuthRepo) FindUser(ctx context.Context, email string) (model.Users, error) {

	sql := "SELECT id, name, email, password, role from users WHERE email=$1"
	args := []any{email}

	var data model.Users
	if err := ar.db.QueryRow(ctx, sql, args...).Scan(&data.Id, &data.Name, &data.Email, &data.Password, &data.Role); err != nil {
		return model.Users{}, err
	}

	return data, nil
}

func (ar *AuthRepo) NewCreateUser(ctx context.Context, body model.Users) error {

	sql := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"
	args := []any{body.Name, body.Email, body.Password}

	cmd, err := ar.db.Exec(ctx, sql, args...)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}

// FindUserByID dipakai untuk change-password, karena user sudah login
func (ar *AuthRepo) FindUserByID(ctx context.Context, id int) (model.Users, error) {
	sql := "SELECT id, name, email, password, role from users WHERE id=$1"
	args := []any{id}

	var data model.Users
	if err := ar.db.QueryRow(ctx, sql, args...).Scan(&data.Id, &data.Name, &data.Email, &data.Password, &data.Role); err != nil {
		return model.Users{}, err
	}

	return data, nil
}

func (ar *AuthRepo) UpdatePasswordByEmail(ctx context.Context, email, hashedPwd string) error {
	sql := "UPDATE users SET password=$1 WHERE email=$2"
	args := []any{hashedPwd, email}

	cmd, err := ar.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return apperror.ErrNotFound
	}

	return nil
}

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrCredentialsEmpty   = errors.New("email or password cannot be empty")

	// Generic data access related
	ErrNotFound       = errors.New("record not found")
	ErrNoRowsAffected = errors.New("no rows affected")

	// General
	ErrInternal     = errors.New("internal server error")
	ErrUnauthorized = errors.New("unauthorized")
)
