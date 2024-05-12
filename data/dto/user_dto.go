package dto

type UserDTO struct {
	Username string `json:"username" validate:"required,min=5"`
	Email    string `json:"email" validate:"required,email,min=5"`
	Password string `json:"password" validate:"required,min=8"`
	Salt     string `json:"salt" validate:"required,min=8"`
	BaseDto
}
