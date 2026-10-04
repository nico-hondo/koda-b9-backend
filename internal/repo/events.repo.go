package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nico-hondo/internal/dto"
	"github.com/nico-hondo/internal/model"
)

type EventsRepo struct {
	db *pgxpool.Pool
}

func NewEventsRepo(db *pgxpool.Pool) *EventsRepo {
	return &EventsRepo{
		db: db,
	}
}

func (er *EventsRepo) GetEventsRepo(ctx context.Context, filter dto.EventFilterParam) ([]dto.EventsResponse, error) {
	sql := `SELECT e.id, e.title, e.description, e.category, e.location, e.image_url, e.start_time, e.end_time, e.max_participants, e.status, e.community_id, e.organizer_id, COALESCE(ep.total_participants, 0) AS total_participants from events e LEFT JOIN(SELECT event_id, COUNT(user_id) AS total_participants FROM event_participants GROUP BY event_id) ep ON e.id = ep.event_id WHERE 1 = 1`

	var args []interface{}
	argIdx := 1

	// Filter Search Keyword (Title atau Description)
	if filter.Search != "" {
		sql += fmt.Sprintf(" AND (e.title ILIKE $%d OR e.description ILIKE $%d)", argIdx, argIdx)
		args = append(args, "%"+filter.Search+"%")
		argIdx++
	}

	// Filter Category
	if filter.Category != "" && strings.ToLower(filter.Category) != "all" {
		sql += fmt.Sprintf(" AND e.category ILIKE $%d", argIdx)
		args = append(args, filter.Category)
		argIdx++
	}

	// Filter Location
	if filter.Location != "" && strings.ToLower(filter.Location) != "all locations" {
		sql += fmt.Sprintf(" AND e.location ILIKE $%d", argIdx)
		args = append(args, filter.Location)
		argIdx++
	}

	// Sorting
	switch filter.Sort {
	case "popular":
		sql += " ORDER BY total_participants DESC"
	case "almost_full":
		sql += " ORDER BY (e.max_participants - COALESCE(ep.total_participants, 0)) ASC"
	case "recently_added":
		sql += " ORDER BY e.created_at DESC"
	default: // 'upcoming'
		sql += " ORDER BY e.start_time ASC"
	}

	events, err := er.db.Query(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	defer events.Close()

	var data []dto.EventsResponse
	for events.Next() {
		var event dto.EventsResponse
		if err := events.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.Category,
			&event.Location,
			&event.ImageUrl,
			&event.StartTime,
			&event.EndTime,
			&event.MaxParticipants,
			&event.Status,
			&event.CommunityID,
			&event.OrganizerID,
			&event.TotalParticipants,
		); err != nil {
			return nil, err
		}

		data = append(data, event)
	}
	return data, nil
}

func (er *EventsRepo) GetEventById(ctx context.Context, id int) (dto.EventDetailResponse, error) {
	sql := `
		SELECT 
			e.id, e.title, e.description, e.category, e.location, e.image_url, 
			e.start_time, e.end_time, e.max_participants, e.status, e.created_at, e.updated_at, 
			e.community_id, e.organizer_id,
			COALESCE(ep.total_participants, 0) AS total_participants,
			u.name AS organizer_name, COALESCE(u.avatar_url, '') AS organizer_avatar,
			c.name AS community_name
		FROM events e
		JOIN users u ON e.organizer_id = u.id
		JOIN communities c ON e.community_id = c.id
		LEFT JOIN (
			SELECT event_id, COUNT(user_id) AS total_participants 
			FROM event_participants 
			GROUP BY event_id
		) ep ON e.id = ep.event_id
		WHERE e.id = $1
	`
	var evDetail dto.EventDetailResponse

	err := er.db.QueryRow(ctx, sql, id).Scan(&evDetail.ID, &evDetail.Title, &evDetail.Description, &evDetail.Category, &evDetail.Location, &evDetail.ImageUrl,
		&evDetail.StartTime, &evDetail.EndTime, &evDetail.MaxParticipants, &evDetail.Status, &evDetail.CreatedAt, &evDetail.UpdatedAt,
		&evDetail.CommunityID, &evDetail.OrganizerID,
		&evDetail.TotalParticipants,
		&evDetail.Organizer.Name, &evDetail.Organizer.AvatarUrl,
		&evDetail.Organizer.CommunityName)

	evDetail.Organizer.UserID = evDetail.OrganizerID
	evDetail.Organizer.CommunityID = evDetail.CommunityID

	return evDetail, err
}

// ambil speakers bersamaan dengan get eventDetail by query params. Jadi eventId dimasukkan melalui param url
func (er *EventsRepo) GetSpeakersByEventID(ctx context.Context, eventID int) ([]model.Speakers, error) {
	query := `
		SELECT s.id, s.name, s.role, s.work_at
		FROM speakers s
		JOIN event_speakers es ON s.id = es.speaker_id
		WHERE es.event_id = $1
	`

	rows, err := er.db.Query(ctx, query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var speakers []model.Speakers
	for rows.Next() {
		var s model.Speakers
		if err := rows.Scan(&s.Id, &s.Name, &s.Role, &s.Work_At); err != nil {
			return nil, err
		}
		speakers = append(speakers, s)
	}

	return speakers, nil
}

func (er *EventsRepo) JoinEventRepo(ctx context.Context, userId, eventId int) (dto.JoinEvent, error) {
	// 1. Ambil data event dan cek apakah user sudah terdaftar
	checkJoin := `
		SELECT 
			e.title, 
			e.start_time, 
			(ep.user_id IS NOT NULL) AS is_joined
		FROM events e
		LEFT JOIN event_participants ep 
			ON e.id = ep.event_id AND ep.user_id = $2
		WHERE e.id = $1
	`
	var result dto.JoinEvent
	// Pastikan urutan argumen: $1 = eventId, $2 = userId
	err := er.db.QueryRow(ctx, checkJoin, eventId, userId).Scan(
		&result.EventTitle,
		&result.StartTime,
		&result.IsJoined,
	)

	if err != nil {
		return dto.JoinEvent{}, err
	}

	// 2. Jika user SUDAH terdaftar -> Lakukan UNJOIN
	if result.IsJoined {
		delSql := `DELETE FROM event_participants WHERE event_id = $1 AND user_id = $2`
		_, err := er.db.Exec(ctx, delSql, eventId, userId)
		if err != nil {
			return dto.JoinEvent{}, err
		}

		// Status terbaru sekarang adalah FALSE (Unjoined)
		result.IsJoined = false
		return result, nil
	}

	// 3. Jika user BELUM terdaftar -> Lakukan JOIN
	insSql := `INSERT INTO event_participants(event_id, user_id, joined_at) VALUES($1, $2, NOW())`
	_, err = er.db.Exec(ctx, insSql, eventId, userId)
	if err != nil {
		return dto.JoinEvent{}, err
	}

	// Status terbaru sekarang adalah TRUE (Joined)
	result.IsJoined = true
	return result, nil
}

func (er *EventsRepo) GetUpcomingEventRepo(ctx context.Context) ([]dto.EventsResponse, error) {
	sql := `
		SELECT 
			e.id, 
			e.title, 
			e.description, 
			e.category, 
			e.location, 
			e.image_url, 
			e.start_time, 
			e.end_time, 
			e.max_participants, 
			e.status, 
			e.community_id, 
			e.organizer_id, 
			COALESCE(ep.total_participants, 0) AS total_participants
		FROM events e 
		LEFT JOIN (
			SELECT event_id, COUNT(user_id) AS total_participants 
			FROM event_participants 
			GROUP BY event_id
		) ep ON e.id = ep.event_id 
		WHERE e.start_time >= NOW()
		ORDER BY e.start_time ASC
		LIMIT 1
	`

	rows, err := er.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []dto.EventsResponse
	for rows.Next() {
		var ev dto.EventsResponse
		err := rows.Scan(
			&ev.ID,
			&ev.Title,
			&ev.Description,
			&ev.Category,
			&ev.Location,
			&ev.ImageUrl,
			&ev.StartTime,
			&ev.EndTime,
			&ev.MaxParticipants,
			&ev.Status,
			&ev.CommunityID,
			&ev.OrganizerID,
			&ev.TotalParticipants,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, ev)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func (er *EventsRepo) MyEventRepo(ctx context.Context, userId int) ([]dto.EventsResponse, error) {
	sql := `
		SELECT 
    		e.id, e.title, e.description, e.category, e.location, e.image_url, e.start_time, e.end_time, e.max_participants, e.status, e.community_id, e.organizer_id, COALESCE(ep.total_participants, 0) AS total_participants 
		from events e 
		LEFT JOIN(SELECT event_id, user_id, COUNT(user_id) AS total_participants FROM event_participants GROUP BY event_id, user_id) ep ON e.id = ep.event_id 
		WHERE ep.user_id=$1
	`

	rows, err := er.db.Query(ctx, sql, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var myEvents []dto.EventsResponse
	for rows.Next() {
		var ev dto.EventsResponse
		err := rows.Scan(
			&ev.ID,
			&ev.Title,
			&ev.Description,
			&ev.Category,
			&ev.Location,
			&ev.ImageUrl,
			&ev.StartTime,
			&ev.EndTime,
			&ev.MaxParticipants,
			&ev.Status,
			&ev.CommunityID,
			&ev.OrganizerID,
			&ev.TotalParticipants,
		)
		if err != nil {
			return nil, err
		}
		myEvents = append(myEvents, ev)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return myEvents, nil
}
