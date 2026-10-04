package model

type Speakers struct {
	Id      int    `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Role    string `db:"role" json:"role"`
	Work_At string `db:"work_at" json:"work_at"`
}
