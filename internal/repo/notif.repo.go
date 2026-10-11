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
	sql := `
			select 
				n.id, n.user_id, n.title, n.message, n.read_at, nt.id, nt.name, nt.icon, nt.bg_color, nt.icon_color
			from notifications n
			join notification_type nt on n.type_id = nt.id
			where n.user_id = $1`
	args := []any{userId}

	notifs, err := nr.db.Query(ctx, sql, args...)

	if err != nil {
		return []model.Notif{}, err
	}

	defer notifs.Close()

	var data []model.Notif
	for notifs.Next() {
		var notif model.Notif

		if err := notifs.Scan(&notif.Id, &notif.User_Id, &notif.Title, &notif.Message, &notif.Read_At, &notif.Type_Id, &notif.Name, &notif.Icon, &notif.Bg_Color, &notif.Icon_Color); err != nil {
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
