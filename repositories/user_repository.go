package repositories

import (
	"github.com/eleynes/MyK3y/db"

	M "github.com/eleynes/MyK3y/models"
)

func CreateUser(user M.User) (string, error) {
	db := db.AppDB.PostgresDB
	result := db.Create(&user)
	if result.Error != nil {
		return user.ID.String(), result.Error
	}
	return user.ID.String(), nil
}

func GetUserByEmail(email string) (M.User, error) {
	db := db.AppDB.PostgresDB
	var user M.User
	err := db.Where("email = ?", email).First(&user).Error

	return user, err
}

func GetUserByID(id string) (M.User, error) {
	db := db.AppDB.PostgresDB
	var user M.User
	err := db.Where("id = ?", id).First(&user).Error

	return user, err
}
