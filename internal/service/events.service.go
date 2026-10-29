package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/nico-hondo/internal/dto"
	"github.com/nico-hondo/internal/model"
	"github.com/nico-hondo/internal/repo"
	"github.com/redis/go-redis/v9"
)

type EventsService struct {
	er  *repo.EventsRepo
	nr  *repo.NotifRepo
	rdb *redis.Client
}

func NewEventsService(er *repo.EventsRepo, nr *repo.NotifRepo, rdb *redis.Client) *EventsService {
	return &EventsService{
		er:  er,
		nr:  nr,
		rdb: rdb,
	}
}

func buildEventsCacheKey(f dto.EventFilterParam) string {
	return fmt.Sprintf("nicohondo:events:search=%s:category=%s:location=%s:sort=%s",
		strings.ToLower(strings.TrimSpace(f.Search)),
		strings.ToLower(strings.TrimSpace(f.Category)),
		strings.ToLower(strings.TrimSpace(f.Location)),
		f.Sort,
	)
}

func (es *EventsService) GetEventsService(ctx context.Context, filter dto.EventFilterParam) ([]dto.EventsResponse, error) {
	if filter.Sort == "" {
		filter.Sort = "upcoming"
	}

	key := buildEventsCacheKey(filter)

	if val, err := es.rdb.Get(ctx, key).Result(); err != nil {
		if errors.Is(err, redis.Nil) {
			log.Println("key does not exist")
		} else {
			log.Println(err.Error())
		}
	} else {
		var events []dto.EventsResponse
		if err := json.Unmarshal([]byte(val), &events); err != nil {
			log.Println("parse error\nreason: ", err.Error())
		} else {
			return events, nil
		}
	}

	if filter.Sort == "" {
		filter.Sort = "upcoming"
	}

	events, err := es.er.GetEventsRepo(ctx, filter)
	if err != nil {
		return nil, err
	}

	if str, err := json.Marshal(events); err != nil {
		log.Println("stringify error\nreason: ", err.Error())
	} else {
		if err := es.rdb.Set(ctx, key, string(str), 10*time.Minute).Err(); err != nil {
			log.Println("redis set error\nreason: ", err.Error())
		}
	}

	return events, nil
}

func (es *EventsService) GetEventDetailService(ctx context.Context, id int) (dto.EventDetailResponse, error) {
	eventDetail, err := es.er.GetEventById(ctx, id)

	if err != nil {
		return dto.EventDetailResponse{}, err
	}

	speakers, err := es.er.GetSpeakersByEventID(ctx, id)
	if err != nil {
		return dto.EventDetailResponse{}, err
	}

	eventDetail.Speakers = speakers
	if eventDetail.Speakers == nil {
		eventDetail.Speakers = []model.Speakers{} // Mencegah return null di JSON
	}

	return eventDetail, nil
}

func (es *EventsService) JoinEventsService(ctx context.Context, userId, eventId int) (bool, string, error) {
	if userId <= 0 {
		return false, "", errors.New("id user tidak valid")
	}

	if eventId <= 0 {
		return false, "", errors.New("id event tidak valid")
	}

	result, err := es.er.JoinEventRepo(ctx, userId, eventId)
	if err != nil {
		return false, "", err
	}

	// Cek status TERBARU setelah aksi
	if result.IsJoined {
		// Kirim notifikasi jika status terbarunya adalah BERHASIL JOIN
		if err := es.nr.NewNotif(ctx, userId, 1, "Join Event Confirmed", fmt.Sprintf("Hooray, You're registered for %s on %s", result.EventTitle, result.StartTime)); err != nil {
			log.Println("failed join event notif: ", err)
		}
		return true, "Berhasil join event!", nil
	}

	// Jika status terbarunya FALSE, berarti aksi yang terjadi adalah UNJOIN
	if err := es.nr.NewNotif(ctx, userId, 1, "UnJoin Event has been Successful", fmt.Sprintf("It's a shame you left the %s event on %s.", result.EventTitle, result.StartTime)); err != nil {
		log.Println("failed unjoin event notif: ", err)
	}
	return false, "Berhasil unjoin event!", nil
}

// func (es *EventsService) GetUpcomingEventService(ctx context.Context, community_id int) ([]dto.EventsResponse, error) {
// 	events, err := es.er.GetUpcomingEventRepo(ctx, community_id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if events == nil {
// 		events = []dto.EventsResponse{}
// 	}

// 	return events, nil
// }

func (es *EventsService) GetMyEventsService(ctx context.Context, userId int) ([]dto.EventsResponse, error) {
	myEvents, err := es.er.MyEventRepo(ctx, userId)
	if err != nil {
		return nil, err
	}

	if myEvents == nil {
		myEvents = []dto.EventsResponse{}
	}

	return myEvents, nil
}

func (es *EventsService) ToggleSaveEvent(ctx context.Context, userID, eventID int) (dto.SaveEventResponse, error) {
	isSaved, err := es.er.ToggleSaveEvent(ctx, userID, eventID)
	if err != nil {
		return dto.SaveEventResponse{}, err
	}

	return dto.SaveEventResponse{
		EventID: eventID,
		IsSaved: isSaved,
	}, nil
}

func (es *EventsService) GetSavedEvents(ctx context.Context, userID int) ([]dto.SavedEventResponse, error) {
	return es.er.GetSavedEvents(ctx, userID)
}
