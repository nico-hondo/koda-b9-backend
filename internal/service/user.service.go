package service

import (
	"context"

	"github.com/nico-hondo/internal/dto"
	"github.com/nico-hondo/internal/model"
	"github.com/nico-hondo/internal/repo"
)

type UserService struct {
	ur *repo.UserRepo
}

func NewUserService(ur *repo.UserRepo) *UserService {
	return &UserService{
		ur: ur,
	}
}

func (us *UserService) GetProfileUserService(ctx context.Context, userId int) (model.Account, error) {
	myProfile, err := us.ur.GetProfileUserByID(ctx, userId)

	if err != nil {
		return model.Account{}, err
	}

	return myProfile, nil
}

func (us *UserService) ChangeUserProfile(ctx context.Context, body dto.UpdateProfile, userId int, imgUrl string) error {
	return us.ur.ChangeUserProfile(ctx, body.Name, imgUrl, body.Bio, body.Location, userId)
}
