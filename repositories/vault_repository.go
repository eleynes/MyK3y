package repositories

import (
	"github.com/eleynes/MyK3y/db"

	M "github.com/eleynes/MyK3y/models"
)

func CreateVault(vault M.Vault) (string, error) {
	db := db.AppDB.PostgresDB
	result := db.Create(&vault)
	if result.Error != nil {
		return vault.ID.String(), result.Error
	}
	return vault.ID.String(), nil
}

func GetVaultByUserIdAndName(userid string, name string) (M.Vault, error) {
	db := db.AppDB.PostgresDB
	var vault M.Vault
	err := db.Where("userid = ? AND vaultname = ?", userid, name).First(&vault).Error

	return vault, err
}

func GetVaultByID(id string) (M.Vault, error) {
	db := db.AppDB.PostgresDB
	var vault M.Vault
	err := db.Where("id = ?", id).First(&vault).Error

	return vault, err
}
