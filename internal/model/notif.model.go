package model

type Notif struct {
	Id      int    `db:"ud"`
	User_Id int    `db:"user_id"`
	Type_Id int    `db:"type_id"`
	Title   string `db:"title"`
	Message string `db:"message"`
}
