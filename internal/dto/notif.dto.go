package dto

type Notif struct {
	Id      int    `json:"id"`
	User_Id int    `json:"user_id"`
	Type_Id int    `json:"type_id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}
