package dto

type OrganizerDetail struct {
	UserID        int    `json:"user_id"`
	Name          string `json:"name"`
	AvatarUrl     string `json:"avatar_url"`
	CommunityID   int    `json:"community_id"`
	CommunityName string `json:"community_name"`
}
