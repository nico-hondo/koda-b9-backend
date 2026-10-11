package dto

import "time"

type Notif struct {
	Id         int        `json:"id"`
	User_Id    int        `json:"user_id"`
	Title      string     `json:"title"`
	Message    string     `json:"message"`
	Read_At    *time.Time `json:"read_at"`
	Type_Id    int        `json:"type_id"`
	Name       string     `json:"name"`
	Icon       string     `json:"icon"`
	Bg_Color   string     `json:"bg_color"`
	Icon_Color string     `json:"icon_color"`
}
