package db

import (
	"fmt"
	"log"
	"strconv"

	C "github.com/eleynes/MyK3y/config"
	P "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect function
func PostgresConnect() (db *gorm.DB, err error) {
	p := C.Conf.PostgresPort
	// because our config function returns a string, we are parsing our str to int here
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("Error parsing str to int")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", C.Conf.PostgresHost, C.Conf.PostgresUser, C.Conf.PostgresPassword, C.Conf.PostgresDB, port)
	db, err = gorm.Open(P.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	log.Println("PostgreSQL Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	// log.Println("running migrations")
	// db.AutoMigrate(&model.User{})

	// DB = Dbinstance{
	// 	Db: db,
	// }
	return db, nil
}
