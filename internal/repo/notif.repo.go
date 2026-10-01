package repo

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nico-hondo/internal/model"
)

type GetNotifRepo struct {
	db *pgxpool.Pool
}

func GetNewNotifRepo(db *pgxpool.Pool) *GetNotifRepo {
	return &GetNotifRepo{
		db: db,
	}
}

func (nr *GetNotifRepo) GetNotif(ctx context.Context, userId int) ([]model.Notif, error) {
	sql := "SELECT id, user_id, type_id, title, message FROM notifications WHERE user_id=$1"
	args := []any{userId}

	notifs, err := nr.db.Query(ctx, sql, args...)

	if err != nil {
		return []model.Notif{}, err
	}

	defer notifs.Close()

	var data []model.Notif
	for notifs.Next() {
		var notif model.Notif

		if err := notifs.Scan(&notif.Id, &notif.User_Id, &notif.Type_Id, &notif.Title, &notif.Message); err != nil {
			return nil, err
		}

		data = append(data, notif)
	}

	if notifs.Err() != nil {
		return []model.Notif{}, notifs.Err()
	}
	log.Println(data)
	return data, nil
}
