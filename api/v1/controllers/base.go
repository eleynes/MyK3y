package controllers

import (
	"github.com/gofiber/fiber/v2"

	C "github.com/eleynes/MyK3y/config"
)

func Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "success",
		"version": C.Conf.Version,
		"env":     C.Conf.Environment,
	})
}
