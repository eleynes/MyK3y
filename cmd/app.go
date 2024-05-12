package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/gofiber/swagger"

	R "github.com/eleynes/MyK3y/api/v1/routes"
	_ "github.com/eleynes/MyK3y/docs"
	H "github.com/eleynes/MyK3y/handler"
)

// @title MyK3y (Password Management System - API)
// @version 1.0
// @description Password Management System - API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support(Erickson Leynes)
// @contact.email erickson.leynes@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8082
// @securitydefinitions.apikey Authentication
// @in header
// @name Authorization
// @BasePath /
func InitApp() *fiber.App {
	app := fiber.New(
		fiber.Config{
			ErrorHandler: H.ErrorHandler,
		},
	)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, PATCH, HEAD",
	}))

	app.Use(requestid.New())

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "github.com/eleynes/MyK3y/docs/swagger.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
	}))

	R.SetupRoutes(app)

	return app
}
