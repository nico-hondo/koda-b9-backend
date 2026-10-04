package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/nico-hondo/internal/dto"
	"github.com/nico-hondo/internal/model"
	"github.com/nico-hondo/internal/repo"
)

type EventsService struct {
	er *repo.EventsRepo
	nr *repo.NotifRepo
}

func NewEventsService(er *repo.EventsRepo, nr *repo.NotifRepo) *EventsService {
	return &EventsService{
		er: er,
		nr: nr,
	}
}

func (es *EventsService) GetEventsService(ctx context.Context, filter dto.EventFilterParam) ([]dto.EventsResponse, error) {
	if filter.Sort == "" {
		filter.Sort = "upcoming"
	}

	events, err := es.er.GetEventsRepo(ctx, filter)
	if err != nil {
		return nil, err
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

func (es *EventsService) GetUpcomingEventService(ctx context.Context) ([]dto.EventsResponse, error) {
	events, err := es.er.GetUpcomingEventRepo(ctx)
	if err != nil {
		return nil, err
	}

	if events == nil {
		events = []dto.EventsResponse{}
	}

	return events, nil
}

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
