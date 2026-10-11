package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/nico-hondo/internal/dto"
	"github.com/nico-hondo/internal/repo"
	"github.com/redis/go-redis/v9"
)

type NotifService struct {
	nr  *repo.GetNotifRepo
	rdb *redis.Client
}

func NewNotifService(nr *repo.GetNotifRepo, rdb *redis.Client) *NotifService {
	return &NotifService{
		nr:  nr,
		rdb: rdb,
	}
}

func (ns *NotifService) GetAllNotifbyId(ctx context.Context, userId int) ([]dto.Notif, error) {
	key := fmt.Sprintf("nicohondo:notifications:%d", userId)

	// 1. Cek Cache Redis
	if val, err := ns.rdb.Get(ctx, key).Result(); err == nil {
		var cachedNotif []dto.Notif
		if err := json.Unmarshal([]byte(val), &cachedNotif); err == nil && len(cachedNotif) > 0 {
			return cachedNotif, nil
		}
	}

	// 2. Cache Miss -> Ambil dari DB
	notif, err := ns.nr.GetNotif(ctx, userId)
	if err != nil {
		return nil, err // Langsung return jika DB error!
	}

	data := make([]dto.Notif, 0, len(notif))
	for _, val := range notif {
		data = append(data, dto.Notif{
			Id:         val.Id,
			User_Id:    val.User_Id,
			Title:      val.Title,
			Message:    val.Message,
			Read_At:    val.Read_At,
			Type_Id:    val.Type_Id,
			Name:       val.Name,
			Icon:       val.Icon,
			Bg_Color:   val.Bg_Color,
			Icon_Color: val.Icon_Color,
		})
	}

	// 3. Simpan ke Redis HANYA JIKA data ada
	if len(data) > 0 {
		if str, errMarshal := json.Marshal(data); errMarshal == nil {
			_ = ns.rdb.Set(ctx, key, string(str), 10*time.Minute).Err()
		}
	}

	return data, nil
}
