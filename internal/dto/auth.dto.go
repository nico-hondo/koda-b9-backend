package dto

type RegisterRequest struct {
	Name     string `json:"nama" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

type ChangeProfileUser struct {
	Name       string `json:"nama" form:"nama"`
	Avatar_url string `json:"avatar" form:"avatar"`
	Bio        string `json:"bio" form:"bio"`
	Location   string `json:"loc" form:"loc"`
	Job        string `json:"job" form:"job"`
	Workplace  string `json:"workplace" form:"workplace"`
}
