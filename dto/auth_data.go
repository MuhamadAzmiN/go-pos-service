package dto

type UserData struct {
	Fullname string `json:"fullname" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}




