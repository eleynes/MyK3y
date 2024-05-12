package routes

import (
	"github.com/gofiber/fiber/v2"

	C "github.com/eleynes/MyK3y/api/v1/controllers"
	MW "github.com/eleynes/MyK3y/api/v1/middleware"
)

func SetupRoutes(app *fiber.App) {
	// health check
	app.Get("/", C.Health)

	// auth
	auth := app.Group("/auth")
	SetupAuthRoutes(auth)

	// protected
	v1APIVault := app.Group("/api/v1/vault", MW.Protected())
	SetupVaultRoutes(v1APIVault)

	v1APIItem := app.Group("/api/v1/item", MW.Protected())
	SetupItemRoutes(v1APIItem)

	v1APIPassword := app.Group("/api/v1/password", MW.Protected())
	SetupPasswordRoutes(v1APIPassword)
}
