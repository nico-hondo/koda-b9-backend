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

type NotifRepo struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func NewNotifRepo(db *pgxpool.Pool) *NotifRepo {
	return &NotifRepo{
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

func (ar *AuthRepo) NewCreateUser(ctx context.Context, body model.Users) (int, error) {

	sql := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	args := []any{body.Name, body.Email, body.Password}

	var newId int
	err := ar.db.QueryRow(ctx, sql, args...).Scan(&newId)

	if err != nil {
		return 0, err
	}

	return newId, nil
}

func (nr *NotifRepo) NewNotif(ctx context.Context, userId, typeId int, title, message string) error {
	sql := "INSERT INTO notifications (user_id, type_id, title, message) VALUES ($1, $2, $3, $4)"
	args := []any{userId, typeId, title, message}
	_, err := nr.db.Exec(ctx, sql, args...)

	return err
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
	sql := "UPDATE users SET password=$1, updated_at=NOW() WHERE email=$2"
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
