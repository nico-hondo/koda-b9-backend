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
	sql := `SELECT 
				e.id, e.title, e.description, e.category, e.location, e.image_url, 
				e.start_time, e.end_time, e.max_participants, e.status, e.community_id, 
				e.organizer_id, COALESCE(ep.total_participants, 0) AS total_participants 
			FROM events e 
			LEFT JOIN (
				SELECT event_id, COUNT(user_id) AS total_participants 
				FROM event_participants 
				GROUP BY event_id
			) ep ON e.id = ep.event_id 
			WHERE 1 = 1`

	var args []interface{}
	argIdx := 1

	// 1. Filter Search Keyword
	if strings.TrimSpace(filter.Search) != "" {
		sql += fmt.Sprintf(" AND (e.title ILIKE $%d OR e.description ILIKE $%d)", argIdx, argIdx+1)
		searchPattern := "%" + strings.TrimSpace(filter.Search) + "%"
		args = append(args, searchPattern, searchPattern)
		argIdx += 2
	}

	// 2. Filter Category
	if filter.Category != "" && strings.ToLower(filter.Category) != "all" {
		sql += fmt.Sprintf(" AND e.category ILIKE $%d", argIdx)
		args = append(args, "%"+strings.TrimSpace(filter.Category)+"%")
		argIdx++
	}

	// 3. Filter Location
	if filter.Location != "" && strings.ToLower(filter.Location) != "all locations" {
		sql += fmt.Sprintf(" AND e.location ILIKE $%d", argIdx)
		args = append(args, "%"+strings.TrimSpace(filter.Location)+"%")
		argIdx++
	}

	// 4. Sorting
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

	// Debugging: Print SQL query dan argument di console backend untuk memastikan query SQL benar
	// fmt.Printf("SQL QUERY: %s\nARGS: %v\n", sql, args)

	rows, err := er.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []dto.EventsResponse
	for rows.Next() {
		var event dto.EventsResponse
		if err := rows.Scan(
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

func (er *EventsRepo) MyEventRepo(ctx context.Context, userId int) ([]dto.EventsResponse, error) {
	sql := `
		SELECT 
			e.id, e.title, e.description, e.category, e.location, e.image_url, 
			e.start_time, e.end_time, e.max_participants, e.status, e.community_id, 
			e.organizer_id,
			COALESCE(COUNT(ep_all.user_id), 0) AS total_participants
		FROM events e
		INNER JOIN event_participants ep_user ON e.id = ep_user.event_id AND ep_user.user_id = $1
		LEFT JOIN event_participants ep_all ON e.id = ep_all.event_id
		GROUP BY e.id
	`

	rows, err := er.db.Query(ctx, sql, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var myEvents []dto.EventsResponse = []dto.EventsResponse{} // Inisialisasi slice kosong biar tidak return nil
	for rows.Next() {
		var ev dto.EventsResponse
		err := rows.Scan(
			&ev.ID, &ev.Title, &ev.Description, &ev.Category, &ev.Location,
			&ev.ImageUrl, &ev.StartTime, &ev.EndTime, &ev.MaxParticipants,
			&ev.Status, &ev.CommunityID, &ev.OrganizerID, &ev.TotalParticipants,
		)
		if err != nil {
			return nil, err
		}
		myEvents = append(myEvents, ev)
	}

	return myEvents, nil
}

func (e *EventsRepo) ToggleSaveEvent(ctx context.Context, userID, eventID int) (bool, error) {
	checkQuery := `SELECT EXISTS(SELECT 1 FROM event_saved WHERE user_id = $1 AND event_id = $2)`
	var isSaved bool
	err := e.db.QueryRow(ctx, checkQuery, userID, eventID).Scan(&isSaved)
	if err != nil {
		return false, err
	}

	if isSaved {
		deleteQuery := `DELETE FROM event_saved WHERE user_id = $1 AND event_id = $2`
		_, err := e.db.Exec(ctx, deleteQuery, userID, eventID)
		if err != nil {
			return false, err
		}
		return false, nil // Berhasil di-unsave
	}

	insertQuery := `INSERT INTO event_saved (user_id, event_id) VALUES ($1, $2)`
	_, err = e.db.Exec(ctx, insertQuery, userID, eventID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (e *EventsRepo) GetSavedEvents(ctx context.Context, userID int) ([]dto.SavedEventResponse, error) {
	query := `
		SELECT es.user_id, es.event_id, es.saved_at, e.title, e.category, e.location, e.image_url, e.start_time, e.max_participants, COALESCE(count(ep.user_id),0) as total_participants
		FROM event_saved es
		join events e on es.event_id = e.id
		join event_participants ep on es.event_id = ep.event_id 
		WHERE es.user_id = $1
		group by 
			es.user_id, 
			es.event_id, 
			es.saved_at, 
			e.title, 
			e.category, 
			e.location, 
			e.image_url, 
			e.start_time, 
			e.max_participants;
	`
	rows, err := e.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	savedEvents := make([]dto.SavedEventResponse, 0)
	for rows.Next() {
		var item dto.SavedEventResponse
		if err := rows.Scan(&item.UserID, &item.EventID, &item.SavedAt, &item.Title, &item.Category, &item.Location, &item.ImageURL, &item.StartTime, &item.MaxParticipants, &item.TotalParticipants); err != nil {
			return nil, err
		}
		savedEvents = append(savedEvents, item)
	}

	return savedEvents, nil
}
