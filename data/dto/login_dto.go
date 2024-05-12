package dto

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email,min=5"`
	Password string `json:"password" validate:"required,min=8"`
}
