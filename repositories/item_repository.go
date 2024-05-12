package repositories

import (
	"github.com/eleynes/MyK3y/db"

	M "github.com/eleynes/MyK3y/models"
)

func CreateItem(item M.Item) (string, error) {
	db := db.AppDB.PostgresDB
	result := db.Create(&item)
	if result.Error != nil {
		return item.ID.String(), result.Error
	}
	return item.ID.String(), nil
}

// func GetVaultByUserIdAndName(userid string, name string) (M.Vault, error) {
// 	db := db.DB.Db
// 	var vault M.Vault
// 	err := db.Where(&vault, "userid = ? AND vaultname = ?", userid, name).Error

// 	return vault, err
// }

func GetItemByID(id string) (M.Item, error) {
	db := db.AppDB.PostgresDB
	var item M.Item
	err := db.Where("id = ?", id).First(&item).Error

	return item, err
}
