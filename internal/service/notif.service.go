package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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

	if val, err := ns.rdb.Get(ctx, key).Result(); err != nil {
		if errors.Is(err, redis.Nil) {
			log.Println("key does not exist")
		} else {
			log.Println(err.Error())
		}
	} else {
		var notif []dto.Notif
		if err := json.Unmarshal([]byte(val), &notif); err != nil {
			log.Println("parse error\nreason: ", err.Error())
		} else {
			return notif, nil
		}
	}

	//cache miss, ambil dari db
	notif, err := ns.nr.GetNotif(ctx, userId)

	data := make([]dto.Notif, 0, len(notif))
	for _, val := range notif {
		data = append(data, dto.Notif{
			Id:      val.Id,
			User_Id: val.User_Id,
			Type_Id: val.Type_Id,
			Title:   val.Title,
			Message: val.Message,
		})
	}
	log.Println(data)

	if str, err := json.Marshal(data); err != nil {
		log.Println("stringify error\nreason: ", err.Error())
	} else {
		if err := ns.rdb.Set(ctx, key, string(str), 10*time.Minute).Err(); err != nil {
			log.Println("redis set error\nreason: ", err.Error())
		}
	}
	return data, err
}
