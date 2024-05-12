package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID                uuid.UUID `json:"ID" gorm:"type:uuid;"`
	Vaultid           uuid.UUID `json:"Vaultid"`
	Username          string    `json:"Username"`
	Encryptedpassword string    `json:"Encryptedpassword"`
	Itemname          string    `json:"Itemname"`
	Url               string    `json:"Url"`
	Notes             string    `json:"Notes"`
}

// Item struct
type Items struct {
	Items []Item `json:"items"`
}

func (item *Item) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	item.ID = uuid.New()
	return
}
