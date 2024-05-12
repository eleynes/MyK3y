package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                 uuid.UUID `json:"ID" gorm:"type:uuid;"`
	Username           string    `json:"Username" gorm:"unique"`
	Email              string    `json:"Email" gorm:"unique"`
	Masterpasswordhash string    `json:"Masterpasswordhash"`
	Masterpasswordsalt string    `json:"Masterpasswordsalt"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.ID = uuid.New()
	return
}
