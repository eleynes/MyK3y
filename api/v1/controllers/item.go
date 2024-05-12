package controllers

import (
	"github.com/gofiber/fiber/v2"

	S "github.com/eleynes/MyK3y/api/v1/services"
	D "github.com/eleynes/MyK3y/data/dto"
	H "github.com/eleynes/MyK3y/handler"
)

// CreateItem   Create Item
// @Summary      Create Item
// @Description  Create Item
// @Tags         Item
// @Accept       json
// @Produce      json
// @Param createItem body D.CreateItemDTO true "Create new Item"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Security 	 Authentication
// @Router       /api/v1/item [post]
func CreateItem(ctx *fiber.Ctx) error {
	createItemDTO := new(D.CreateItemDTO)

	if err := ctx.BodyParser(createItemDTO); err != nil {
		return H.BuildError(ctx, "Invalid body", fiber.StatusBadRequest, err)
	}

	item, serviceErr := S.CreateItem(ctx, createItemDTO)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, "Successfully Created", item)
}

// GetItemByID  Get Item By ID
// @Summary      Get Item By ID
// @Description  Get Item By ID
// @Tags         Item
// @Accept       json
// @Produce      json
// @Param        id    path     string  false  "Item search by ID"  Format(uuid)
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Security 	 Authentication
// @Router       /api/v1/item/{id} [get]
func GetItemByID(ctx *fiber.Ctx) error {
	item, serviceErr := S.GetItemById(ctx)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, "Successfully Fetched", item)
}
