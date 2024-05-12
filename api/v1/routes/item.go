package routes

import (
	"github.com/gofiber/fiber/v2"

	C "github.com/eleynes/MyK3y/api/v1/controllers"
)

func SetupItemRoutes(router fiber.Router) {

	router.Post("/", C.CreateItem)

	router.Get("/:id", C.GetItemByID)

	// router.Patch("/:id", C.UpdateItem)

	// router.Delete("/:id", C.DeleteItem)

}
