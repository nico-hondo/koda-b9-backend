package model

import "time"

type Notif struct {
	Id         int        `db:"id"`
	User_Id    int        `db:"user_id"`
	Title      string     `db:"title"`
	Message    string     `db:"message"`
	Read_At    *time.Time `db:"read_at"`
	Type_Id    int        `db:"type_id"`
	Name       string     `db:"name"`
	Icon       string     `db:"icon"`
	Bg_Color   string     `db:"bg_color"`
	Icon_Color string     `db:"icon_color"`
}
