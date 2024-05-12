package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vault struct {
	gorm.Model
	ID        uuid.UUID `json:"ID" gorm:"type:uuid;"`
	Userid    uuid.UUID `json:"Userid" gorm:"unique"`
	Vaultname string    `json:"Vaultname" gorm:"unique"`
}

// Vault struct
type Vaults struct {
	Vaults []Vault `json:"vaults"`
}

func (vault *Vault) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	vault.ID = uuid.New()
	return
}
