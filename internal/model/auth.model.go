package model

import "time"

type Users struct {
	Id       int    `db:"id"`
	Name     string `db:"nama"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

type Account struct {
	Id        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"nama"`
	Email     string    `json:"email" db:"email"`
	AvatarURL *string   `json:"avatar_url" db:"avatar_url"`
	Bio       *string   `json:"bio" db:"bio"`
	Location  *string   `json:"location" db:"location"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
