package controllers

import (
	"github.com/gofiber/fiber/v2"

	S "github.com/eleynes/MyK3y/api/v1/services"
	D "github.com/eleynes/MyK3y/data/dto"
	H "github.com/eleynes/MyK3y/handler"
)

// CreateVault   Create Vault
// @Summary      Create Vault
// @Description  Create Vault
// @Tags         Vault
// @Accept       json
// @Produce      json
// @Param createvault body D.CreateVaultDTO true "Create new vault"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Security 	 Authentication
// @Router       /api/v1/vault [post]
func CreateVault(ctx *fiber.Ctx) error {
	createVaultDTO := new(D.CreateVaultDTO)

	if err := ctx.BodyParser(createVaultDTO); err != nil {
		return H.BuildError(ctx, "Invalid body", fiber.StatusBadRequest, err)
	}

	vault, serviceErr := S.CreateVault(ctx, createVaultDTO)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, "Successfully Created", vault)
}

// GetVaultByID  Get Vault By ID
// @Summary      Get Vault By ID
// @Description  Get Vault By ID
// @Tags         Vault
// @Accept       json
// @Produce      json
// @Param        id    path     string  false  "vault search by ID"  Format(uuid)
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Security 	 Authentication
// @Router       /api/v1/vault/{id} [get]
func GetVaultByID(ctx *fiber.Ctx) error {
	vault, serviceErr := S.GetVaultById(ctx)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, "Successfully Fetched", vault)
}
