package service

import (
	"context"
	"errors"
	"log"

	"github.com/nico-hondo/internal/dto"
	apperror "github.com/nico-hondo/internal/error"
	"github.com/nico-hondo/internal/repo"
)

type TestimoniService struct {
	tr *repo.GetTestimoniRepo
	nr *repo.NotifRepo
}

func NewTestimoniService(tr *repo.GetTestimoniRepo, nr *repo.NotifRepo) *TestimoniService {
	return &TestimoniService{
		tr: tr,
		nr: nr,
	}
}

func (ts *TestimoniService) GetTestimoniService(ctx context.Context) ([]dto.Testimoni, error) {
	testimonies, err := ts.tr.GetTestimoniRepo(ctx)
	if err != nil {
		return []dto.Testimoni{}, err
	}

	var dtoTestimonies []dto.Testimoni
	for _, t := range testimonies {
		dtoTestimonies = append(dtoTestimonies, dto.Testimoni{
			Id:         t.ID,
			User_Id:    t.User_Id,
			Comment:    t.Comment,
			Created_At: t.Created_At,
			Name:       t.Name,
			Job:        t.Job,
			Workplace:  t.Workplace,
		})
	}

	return dtoTestimonies, nil
}

func (ts *TestimoniService) NewCreateTestimoniService(ctx context.Context, userId int, comment string) error {
	if len(comment) == 0 {
		return errors.New("comment cannot be empty")
	}

	if err := ts.tr.NewCreateTestimoniRepo(ctx, userId, comment); err != nil {
		return apperror.ErrInternal
	}

	if err := ts.nr.NewNotif(ctx, userId, 2, "Testimonial Confirmed", "Congratulations your Testimonial has been successfully created!"); err != nil {
		log.Println("failed create welcome notif: ", err)
	}

	return nil
}
