package dto

type UserData struct {
	Fullname string `json:"fullName" validate:"required"`
	Email    string `gorm:"unique" json:"email" validate:"required,email" `
	Password string `json:"password" validate:"required,min=6"`
}
