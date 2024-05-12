package services

import (
	"time"

	DB "github.com/eleynes/MyK3y/db"
	M "github.com/eleynes/MyK3y/models"
)

func GetCachedData(modelVar interface{}, key string) ([]byte, error) {
	vaultDataBytes, err := DB.AppDB.RedisDB.Get(M.GetType(modelVar) + key)
	return vaultDataBytes, err
}

func SetCacheData(modelVar interface{}, key string, value []byte) error {
	err := DB.AppDB.RedisDB.Set(M.GetType(modelVar)+key, value, 24*time.Hour)

	return err
}

func DeleteCachedData(modelVar interface{}, key string) error {
	err := DB.AppDB.RedisDB.Delete(M.GetType(modelVar) + key)
	return err
}
