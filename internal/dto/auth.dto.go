package dto

type RegisterRequest struct {
	Name     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

type ChangeProfileUser struct {
	Name       string `json:"nama" form:"nama"`
	Avatar_url string `json:"avatar" form:"avatar"`
	Bio        string `json:"bio" form:"bio"`
	Location   string `json:"loc" form:"loc"`
	Job        string `json:"job" form:"job"`
	Workplace  string `json:"workplace" form:"workplace"`
}
