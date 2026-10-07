package model

import "time"

type Testimoni struct {
	ID         int       `json:"id"`
	User_Id    int       `json:"user_id"`
	Comment    string    `json:"comment"`
	Created_At time.Time `json:"created_at"`
	Name       string    `json:"name"`
	Job        *string   `json:"job"`
	Workplace  *string   `json:"workplace"`
}

type NewTestimoni struct {
	User_Id int    `json:"user_id"`
	Comment string `json:"comment"`
}
