package services

import (
	"encoding/hex"
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	D "github.com/eleynes/MyK3y/data/dto"
	M "github.com/eleynes/MyK3y/models"
	R "github.com/eleynes/MyK3y/repositories"
	T "github.com/eleynes/MyK3y/types"
)

func CreateItem(ctx *fiber.Ctx, createItemDTO *D.CreateItemDTO) (*M.Item, *T.ServiceError) {

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	userData, err := R.GetUserByEmail(email)
	if err != nil {
		return nil, &T.ServiceError{
			Message: "Unable to create item",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	vaultData, err := R.GetVaultByUserIdAndName(userData.ID.String(), createItemDTO.Vaultname)
	if err == nil && vaultData.ID == uuid.Nil {
		return nil, &T.ServiceError{
			Message: "Unable to create item",
			Error:   err,
			Code:    fiber.StatusBadRequest,
		}
	}

	key := []byte("16byteAESKey1234")
	message := []byte(createItemDTO.Password)

	encryptedpassword := Encrypt(message, key)

	item := M.Item{
		Vaultid:           vaultData.ID,
		Itemname:          createItemDTO.Itemname,
		Username:          createItemDTO.Username,
		Encryptedpassword: hex.EncodeToString(encryptedpassword),
		Url:               createItemDTO.Url,
		Notes:             createItemDTO.Notes,
	}

	itemId, err := R.CreateItem(item)

	if err != nil {
		if errors.Is(err, gorm.ErrInvalidData) {
			return nil, &T.ServiceError{
				Message: "Invalid Data",
				Error:   err,
				Code:    fiber.StatusNotFound,
			}
		}
		return nil, &T.ServiceError{
			Message: "Unable to create c",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}
	itemData, err := R.GetItemByID(itemId)
	if err != nil {
		return nil, &T.ServiceError{
			Message: "Unable to create item",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	return &itemData, nil
}

func GetItemById(ctx *fiber.Ctx) (*M.Item, *T.ServiceError) {
	itemID := ctx.Params("id")

	itemData, err := R.GetItemByID(itemID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &T.ServiceError{
				Message: "Item not found",
				Error:   err,
				Code:    fiber.StatusNotFound,
			}
		}
		return nil, &T.ServiceError{
			Message: "Unable to get item",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	return &itemData, nil
}
