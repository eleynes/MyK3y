package routes

import (
	"io"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"

	// "github.com/hrshadhin/fiber-go-boilerplate/app/controller"
	// "github.com/hrshadhin/fiber-go-boilerplate/app/model"
	// "github.com/hrshadhin/fiber-go-boilerplate/app/repository"
	// "github.com/hrshadhin/fiber-go-boilerplate/pkg/config"
	// "github.com/hrshadhin/fiber-go-boilerplate/platform/database"

	C "github.com/eleynes/MyK3y/config"
	DB "github.com/eleynes/MyK3y/db"
	M "github.com/eleynes/MyK3y/models"
	R "github.com/eleynes/MyK3y/repositories"

	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublicRoutesLogin(t *testing.T) {
	setUpTPuR()
	defer tearDownTPuR()

	app := fiber.New()
	auth := app.Group("/auth")
	SetupAuthRoutes(auth)

	postURLTests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
		requestBody   io.Reader
	}{
		{
			description:   "Login (success)",
			route:         "/auth/login",
			expectedError: false,
			expectedCode:  200,
			requestBody:   strings.NewReader(`{"email": "test1@domain.com", "password": "Password123"}`),
		},
		{
			description:   "Login (invalid password)",
			route:         "/auth/login",
			expectedError: false,
			expectedCode:  401,
			requestBody:   strings.NewReader(`{"email": "test2@domain.com", "password": "WrongPassword123"}`),
		},
		{
			description:   "Login (invalid email)",
			route:         "/auth/login",
			expectedError: false,
			expectedCode:  401,
			requestBody:   strings.NewReader(`{"email": "test3@domain.com", "password": "Password123"}`),
		},
		{
			description:   "Register (success)",
			route:         "/auth/register",
			expectedError: false,
			expectedCode:  200,
			requestBody: strings.NewReader(`{
				"username" : "test3",
				"email":"test3@domain.com",
				"password" : "Password123",
				"salt" : "MySalt123"
			}`),
		},
		{
			description:   "Register (email taken)",
			route:         "/auth/register",
			expectedError: false,
			expectedCode:  400,
			requestBody: strings.NewReader(`{
				"username" : "test1",
				"email":"test1@domain.com",
				"password" : "Password123",
				"salt" : "MySalt123"
			}`),
		},
	}

	for _, test := range postURLTests {
		req := httptest.NewRequest("POST", test.route, test.requestBody)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description+req.Host)

		// As expected errors lead to broken responses,
		// the next test case needs to be processed.
		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func setUpTPuR() {

	if godotenv.Load("../../../.env.test") != nil {
		log.Fatal("Error loading .env file")
	}

	_, configErr := C.New()

	if configErr != nil {
		log.Fatal(configErr)
	}

	postgresDB, postgresDBErr := DB.PostgresConnect()
	redisDB, redisDBErr := DB.RedisConnect()

	if postgresDBErr != nil {
		panic(postgresDBErr)
	}

	if redisDBErr != nil {
		panic(redisDBErr)
	}

	DB.AppDB = DB.Dbinstance{
		PostgresDB: postgresDB,
		RedisDB:    redisDB,
	}

	users := []M.User{
		{
			ID:                 uuid.Must(uuid.Parse("91f63a89-fe89-4649-abe4-d6704d69ea83")),
			Username:           "test1",
			Email:              "test1@domain.com",
			Masterpasswordhash: "$2a$10$XAZ04f3pFd6n0AnZBtZ8WuUpP7KnBNXr/rlHVoXP.tZFHdhKxMXUm",
			Masterpasswordsalt: "MySalt123",
		},
		{
			ID:                 uuid.Must(uuid.Parse("91f63a89-fe89-4649-abe4-d6704d69ea82")),
			Username:           "test2",
			Email:              "test2@domain.com",
			Masterpasswordhash: "$2a$10$XAZ04f3pFd6n0AnZBtZ8WuUpP7KnBNXr/rlHVoXP.tZFHdhKxMXUm",
			Masterpasswordsalt: "MySalt123",
		},
	}

	for _, user := range users {
		if _, err := R.CreateUser(user); err != nil {
			panic(err)
		}
	}
}

func tearDownTPuR() {
	err := DB.AppDB.PostgresDB.Exec(`TRUNCATE TABLE "users" CASCADE;`).Error
	if err != nil {
		panic(err)
	}

}
