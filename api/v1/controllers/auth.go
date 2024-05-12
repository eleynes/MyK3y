package controllers

import (
	"github.com/gofiber/fiber/v2"

	S "github.com/eleynes/MyK3y/api/v1/services"
	D "github.com/eleynes/MyK3y/data/dto"
	H "github.com/eleynes/MyK3y/handler"
)

// Register      User Registration
// @Summary      User Registration
// @Description  User Registration
// @Tags         User
// @Accept       json
// @Produce      json
// @Param UserRegistration body D.UserDTO true "Register new User"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Security 	 Authentication
// @Router       /auth/register [post]
func Register(ctx *fiber.Ctx) error {
	userDTO := new(D.UserDTO)

	if err := ctx.BodyParser(userDTO); err != nil {
		return H.BuildError(ctx, "Invalid body", fiber.StatusBadRequest, err)
	}

	user, serviceErr := S.Register(userDTO)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, "Successfully Registered", user)
}

// Login         User Login
// @Summary      User Login
// @Description  User Login
// @Tags         User
// @Accept       json
// @Produce      json
// @Param UserLogin body D.LoginDTO true "User Login"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Security 	 Authentication
// @Router       /auth/login [post]
func Login(ctx *fiber.Ctx) error {
	loginDTO := new(D.LoginDTO)

	if err := ctx.BodyParser(loginDTO); err != nil {
		return H.BuildError(ctx, "Invalid body", fiber.StatusBadRequest, err)
	}
	token, serviceErr := S.Login(loginDTO)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, "Successfully Login", token)
}
