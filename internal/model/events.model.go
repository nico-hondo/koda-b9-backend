package model

import "time"

type Events struct {
	ID              int       `db:"id" json:"id"`
	Title           string    `db:"title" json:"title"`
	Description     string    `db:"description" json:"description"`
	Category        string    `db:"category" json:"category"`
	Location        string    `db:"location" json:"location"`
	ImageUrl        string    `db:"image_url" json:"image_url"`
	StartTime       time.Time `db:"start_time" json:"start_time"`
	EndTime         time.Time `db:"end_time" json:"end_time"`
	MaxParticipants int       `db:"max_participants" json:"max_participants"`
	Status          string    `db:"status" json:"status"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
	CommunityID     int       `db:"community_id" json:"community_id"`
	OrganizerID     int       `db:"organizer_id" json:"organizer_id"`
}

type JoinEvents struct {
	Event_Id  int       `db:"event_id" json:"event_id"`
	User_Id   int       `db:"user_id" json:"user_id"`
	Joined_At time.Time `db:"joined_at" json:"joined_at"`
}
