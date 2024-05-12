package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/eleynes/MyK3y/config"
	D "github.com/eleynes/MyK3y/data/dto"
	M "github.com/eleynes/MyK3y/models"
	R "github.com/eleynes/MyK3y/repositories"
	T "github.com/eleynes/MyK3y/types"
)

func CreateVault(ctx *fiber.Ctx, createVaultDTO *D.CreateVaultDTO) (*M.Vault, *T.ServiceError) {

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	userData, err := R.GetUserByEmail(email)
	if err != nil {
		return nil, &T.ServiceError{
			Message: "Unable to create vault",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	existingVault, err := R.GetVaultByUserIdAndName(userData.ID.String(), createVaultDTO.Vaultname)
	if err == nil && existingVault.ID != uuid.Nil {
		return nil, &T.ServiceError{
			Message: "Vault Already Existing: " + existingVault.Vaultname,
			Error:   err,
			Code:    fiber.StatusBadRequest,
		}
	}

	vault := M.Vault{
		Userid:    userData.ID,
		Vaultname: createVaultDTO.Vaultname,
	}

	vaultId, err := R.CreateVault(vault)

	if err != nil {
		if errors.Is(err, gorm.ErrInvalidData) {
			return nil, &T.ServiceError{
				Message: "Invalid Data",
				Error:   err,
				Code:    fiber.StatusNotFound,
			}
		}
		return nil, &T.ServiceError{
			Message: "Unable to create vault",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}
	vaultData, err := R.GetVaultByID(vaultId)
	if err != nil {
		return nil, &T.ServiceError{
			Message: "Unable to create vault",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	return &vaultData, nil
}

func GetVaultById(ctx *fiber.Ctx) (*M.Vault, *T.ServiceError) {
	vaultId := ctx.Params("id")
	var vaultData M.Vault
	if config.Conf.EnableCache {
		vaultDataBytes, err := GetCachedData(vaultData, vaultId)
		if err != nil && err != redis.Nil {
			return nil, &T.ServiceError{
				Message: "Unable to return cache vault",
				Error:   err,
				Code:    fiber.StatusInternalServerError,
			}
		}
		if len(vaultDataBytes) > 0 {
			json.Unmarshal(vaultDataBytes, &vaultData)
			fmt.Println("Cached Data", M.GetType(vaultData))
			return &vaultData, nil
		}
	}

	vaultData, err := R.GetVaultByID(vaultId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &T.ServiceError{
				Message: "Vault not found",
				Error:   err,
				Code:    fiber.StatusNotFound,
			}
		}
		return nil, &T.ServiceError{
			Message: "Unable to get vault",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	if config.Conf.EnableCache {
		jsonBytes, _ := json.Marshal(vaultData)
		errSaveToRedis := SetCacheData(vaultData, vaultId, jsonBytes)
		if errSaveToRedis != nil {
			return nil, &T.ServiceError{
				Message: "Unable to cache vault",
				Error:   err,
				Code:    fiber.StatusInternalServerError,
			}
		}
	}

	return &vaultData, nil
}
