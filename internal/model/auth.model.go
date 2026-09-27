package model

type Users struct {
	Id       int    `db:"id"`
	Name     string `db:"nama"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
}
