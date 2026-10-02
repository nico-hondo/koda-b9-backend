package repo

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nico-hondo/internal/model"
)

type GetTestimoniRepo struct {
	db *pgxpool.Pool
}

func NewTestimoniRepo(db *pgxpool.Pool) *GetTestimoniRepo {
	return &GetTestimoniRepo{
		db: db,
	}
}

func (tr *GetTestimoniRepo) GetTestimoniRepo(ctx context.Context) ([]model.Testimoni, error) {
	sql := "select t.id, t.user_id, t.comment, t.created_at, u.name, u.job, u.workplace from testimonies t join users u on t.user_id = u.id;"

	testi, err := tr.db.Query(ctx, sql)

	if err != nil {
		return []model.Testimoni{}, err
	}

	defer testi.Close()

	var data []model.Testimoni
	for testi.Next() {
		var testimoni model.Testimoni

		if err := testi.Scan(&testimoni.ID, &testimoni.User_Id, &testimoni.Comment, &testimoni.Created_At, &testimoni.Name, &testimoni.Job, &testimoni.Workplace); err != nil {
			return nil, err
		}

		data = append(data, testimoni)
	}

	if testi.Err() != nil {
		return []model.Testimoni{}, testi.Err()
	}

	log.Println(data)

	return data, nil
}

func (tr *GetTestimoniRepo) NewCreateTestimoniRepo(ctx context.Context, userId int, comment string) error {
	sql := "INSERT INTO testimonies(user_id, comment, created_at) VALUES ($1, $2, NOW())"
	args := []any{userId, comment}

	_, err := tr.db.Exec(ctx, sql, args...)

	return err
}
