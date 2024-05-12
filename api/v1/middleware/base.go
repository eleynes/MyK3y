package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"

	C "github.com/eleynes/MyK3y/config"
	JW "github.com/gofiber/jwt/v2"
)

var CurrentLoggedInUser string

// Protected protect routes
func Protected() func(*fiber.Ctx) error {
	return JW.New(JW.Config{
		SigningKey:   []byte(C.Conf.Secret),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	log.Println(err.Error())
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})

	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
	}
}
