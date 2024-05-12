package dto

type GeneratePasswordDTO struct {
	Length              int    `json:"length"`
	Count               int    `json:"count"`
	PasswordType        string `json:"passwordType"`
	IsNumbersIncluded   bool   `json:"isNumbersIncluded"`
	IsSymbolsIncluded   bool   `json:"isSymbolsIncluded"`
	IsUppercaseIncluded bool   `json:"isUppercaseIncluded"`
}
