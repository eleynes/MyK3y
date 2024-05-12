package routes

import (
	"github.com/gofiber/fiber/v2"

	C "github.com/eleynes/MyK3y/api/v1/controllers"
)

func SetupAuthRoutes(router fiber.Router) {

	router.Post("/login", C.Login)
	router.Post("/register", C.Register)

}
