package main

import (
	"log"

	"github.com/joho/godotenv"

	CMD "github.com/eleynes/MyK3y/cmd"
	C "github.com/eleynes/MyK3y/config"
	DB "github.com/eleynes/MyK3y/db"
)

func main() {

	if godotenv.Load(".env") != nil {
		log.Fatal("Error loading .env file")
	}

	confVars, configErr := C.New()

	if configErr != nil {
		log.Fatal(configErr)
	}

	postgresDB, postgresDBErr := DB.PostgresConnect()
	redisDB, redisDBErr := DB.RedisConnect()

	if postgresDBErr != nil {
		log.Fatal(postgresDBErr)
	}

	if redisDBErr != nil {
		log.Fatal(redisDBErr)
	}

	DB.AppDB = DB.Dbinstance{
		PostgresDB: postgresDB,
		RedisDB:    redisDB,
	}

	app := CMD.InitApp()

	app.Listen(confVars.Port)
}
