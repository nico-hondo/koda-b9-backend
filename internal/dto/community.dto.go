package dto

import "time"

type CommunityFilterParam struct {
	Search   string `form:"search" json:"search"`
	Category string `form:"category" json:"category"`
	IsJoined string `form:"is_joined" json:"is_joined"` // 'upcoming', 'popular', 'almost_full', 'recently_added'
}

type CommunityResponse struct {
	Id            int       `db:"id" json:"id"`
	Name          string    `db:"name" json:"name"`
	Slug          string    `db:"slug" json:"slug"`
	Description   string    `db:"description" json:"description"`
	Category      string    `db:"category" json:"category"`
	ImageUrl      string    `db:"image_url" json:"image_url"`
	Location      string    `db:"location" json:"location"`
	Is_Active     bool      `db:"is_active" json:"is_active"`
	Created_At    time.Time `db:"created_at" json:"created_at"`
	Total_Members int       `db:"total_members" json:"total_members"`
}

type MemberResponse struct {
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
	Job       string `json:"job"`
	AvatarUrl string `json:"avatar_url"`
}

type CommunityDetailResponse struct {
	ID                  int              `json:"id"`
	Name                string           `json:"name"`
	Slug                string           `json:"slug"`
	Description         string           `json:"description"`
	Category            string           `json:"category"`
	ImageUrl            string           `json:"image_url"`
	Location            string           `json:"location"`
	IsActive            bool             `json:"is_active"`
	CreatedAt           time.Time        `json:"created_at"`
	TotalMembers        int              `json:"total_members"`
	TotalUpcomingEvents int              `json:"total_upcoming_events"`
	IsJoined            bool             `json:"is_joined"`
	Members             []MemberResponse `json:"members"`
}

type PopularCommunityResponse struct {
	Id                  int       `json:"id"`
	Name                string    `json:"name"`
	Slug                string    `json:"slug"`
	Description         string    `json:"description"`
	Category            string    `json:"category"`
	ImageUrl            string    `json:"image_url"`
	Location            string    `json:"location"`
	Is_Active           bool      `json:"is_active"`
	Created_At          time.Time `json:"created_at"`
	Total_Members       int       `json:"total_members"`
	TotalUpcomingEvents int       `json:"total_upcoming_events"`
}
