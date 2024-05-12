package dto

type CreateItemDTO struct {
	Vaultname string `json:"vaultname"`
	Itemname  string `json:"itemname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Url       string `json:"url"`
	Notes     string `json:"notes"`
	BaseDto
}
