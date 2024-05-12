package controllers

import (
	"github.com/gofiber/fiber/v2"

	S "github.com/eleynes/MyK3y/api/v1/services"
	D "github.com/eleynes/MyK3y/data/dto"
	H "github.com/eleynes/MyK3y/handler"
)

// GeneratePassword   Generate Password
// @Summary      Generate Password
// @Description  Generate Password
// @Tags         Password
// @Accept       json
// @Produce      json
// @Param GeneratePassword body D.GeneratePasswordDTO true "Generate new Password"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Security 	 Authentication
// @Router       /api/v1/password [post]
func GeneratePassword(ctx *fiber.Ctx) error {
	generatePasswordDTO := new(D.GeneratePasswordDTO)

	if err := ctx.BodyParser(generatePasswordDTO); err != nil {
		return H.BuildError(ctx, "Invalid body", fiber.StatusBadRequest, err)
	}

	passwords, serviceErr := S.GeneratePassword(ctx, generatePasswordDTO)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, "Successfully Generated", passwords)
}
