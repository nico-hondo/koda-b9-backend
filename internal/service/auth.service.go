package service

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/nico-hondo/internal/dto"
	apperror "github.com/nico-hondo/internal/error"
	"github.com/nico-hondo/internal/model"
	"github.com/nico-hondo/internal/repo"
	"github.com/nico-hondo/pkg"
)

type AuthService struct {
	ar *repo.AuthRepo
}

func NewAuthService(ar *repo.AuthRepo) *AuthService {
	return &AuthService{
		ar: ar,
	}
}

func (as *AuthService) CreateUser(ctx context.Context, body dto.RegisterRequest) error {
	//validasi
	if len(body.Name) == 0 || len(body.Email) == 0 || len(body.Password) == 0 {
		return errors.New("name, email and password tidak boleh kosong")
	}
	//Cek exists di repo
	_, err := as.ar.FindUser(ctx, body.Email)
	if err == nil {
		return apperror.ErrEmailAlreadyExists
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	//hash pwd
	hc := pkg.NewRecommendedHashConfig()
	hashedPwd := hc.GenHash(body.Password)

	//menambahkan data ke db
	if err := as.ar.NewCreateUser(ctx, model.Users{
		Name:     body.Name,
		Email:    body.Email,
		Password: hashedPwd,
	}); err != nil {
		return err
	}
	return nil
}

func (as *AuthService) Login(ctx context.Context, body dto.LoginRequest) (string, error) {
	//validas
	if len(body.Email) == 0 || len(body.Password) == 0 {
		return "", apperror.ErrCredentialsEmpty
	}

	//cek account apakah sudah ada di db (melalui repo)
	acc, err := as.ar.FindUser(ctx, body.Email)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", apperror.ErrInvalidCredentials
		}
		return "", err
	}

	if err := pkg.Compare(body.Password, acc.Password); err != nil {
		log.Printf("DEBUG compare failed: %v", err)
		return "", apperror.ErrInvalidCredentials
	}

	//Jika tidak terjadi error, maka berikan nilai truenya
	claims := pkg.NewJwtClaims(acc.Id, acc.Role)

	return claims.GenToken()
}

func (as *AuthService) ChangePassword(ctx context.Context, userID int, oldPassword, newPassword string) error {
	if len(oldPassword) == 0 || len(newPassword) == 0 {
		return apperror.ErrCredentialsEmpty
	}

	acc, err := as.ar.FindUserByID(ctx, userID)
	if err != nil {
		return apperror.ErrInternal
	}

	if err := pkg.Compare(oldPassword, acc.Password); err != nil {
		return apperror.ErrInvalidCredentials
	}

	hc := pkg.NewRecommendedHashConfig()
	hashedPwd := hc.GenHash(newPassword)

	if err := as.ar.UpdatePasswordByEmail(ctx, acc.Email, hashedPwd); err != nil {
		return apperror.ErrInternal
	}

	return nil
}
