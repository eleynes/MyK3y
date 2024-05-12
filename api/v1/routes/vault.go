package routes

import (
	"github.com/gofiber/fiber/v2"

	C "github.com/eleynes/MyK3y/api/v1/controllers"
)

func SetupVaultRoutes(router fiber.Router) {

	router.Post("/", C.CreateVault)

	router.Get("/:id", C.GetVaultByID)

}
