package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Password struct {
	ID                uuid.UUID `json:"ID" gorm:"type:uuid;"`
	Encryptedpassword string    `json:"Encryptedpassword"`
	Plainpassword     string    `json:"Plainpassword"`
}

// Password struct
type Passwords struct {
	Items []Password `json:"passwords"`
}

func (password *Password) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	password.ID = uuid.New()
	return
}
