package dto

import (
	"time"

	"github.com/nico-hondo/internal/model"
)

// EventFilterParam menampung query parameter dari URL (misal: /events?search=go&category=technology&sort=popular)
type EventFilterParam struct {
	Search   string `form:"search" json:"search"`
	Category string `form:"category" json:"category"`
	Location string `form:"location" json:"location"`
	Sort     string `form:"sort" json:"sort"` // 'upcoming', 'popular', 'almost_full', 'recently_added'
}

// EventsResponse menampung hasil query gabungan event + total_participants untuk dikirim ke Frontend
type EventsResponse struct {
	ID                int       `db:"id" json:"id"`
	Title             string    `db:"title" json:"title"`
	Description       string    `db:"description" json:"description"`
	Category          string    `db:"category" json:"category"`
	Location          string    `db:"location" json:"location"`
	ImageUrl          string    `db:"image_url" json:"image_url"`
	StartTime         time.Time `db:"start_time" json:"start_time"`
	EndTime           time.Time `db:"end_time" json:"end_time"`
	MaxParticipants   int       `db:"max_participants" json:"max_participants"`
	Status            string    `db:"status" json:"status"`
	CommunityID       int       `db:"community_id" json:"community_id"`
	OrganizerID       int       `db:"organizer_id" json:"organizer_id"`
	TotalParticipants int       `db:"total_participants" json:"total_participants"`
}

type EventDetailResponse struct {
	model.Events
	TotalParticipants int              `json:"total_participants"`
	Organizer         OrganizerDetail  `json:"organizer"`
	Speakers          []model.Speakers `json:"speakers"`
}

type JoinEvent struct {
	IsJoined   bool      `json:"is_joined"`
	EventTitle string    `json:"event_title"`
	StartTime  time.Time `json:"start_time"`
}
