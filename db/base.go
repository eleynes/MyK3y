package db

import (
	"gorm.io/gorm"

	"github.com/gofiber/storage/redis/v3"
	_ "github.com/lib/pq"
)

// Database instance
type Dbinstance struct {
	PostgresDB *gorm.DB
	RedisDB    *redis.Storage
}

var AppDB Dbinstance
