package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	C "github.com/eleynes/MyK3y/config"
	D "github.com/eleynes/MyK3y/data/dto"
	M "github.com/eleynes/MyK3y/models"
	R "github.com/eleynes/MyK3y/repositories"
	T "github.com/eleynes/MyK3y/types"
)

func Register(userDTO *D.UserDTO) (*M.User, *T.ServiceError) {

	password, _ := bcrypt.GenerateFromPassword([]byte(userDTO.Password+userDTO.Salt), bcrypt.DefaultCost)

	user := M.User{
		Username:           userDTO.Username,
		Email:              userDTO.Email,
		Masterpasswordsalt: userDTO.Salt,
		Masterpasswordhash: string(password),
	}

	userId, err := R.CreateUser(user)

	if err != nil {
		if errors.Is(err, gorm.ErrInvalidData) {
			return nil, &T.ServiceError{
				Message: "Invalid Data",
				Error:   err,
				Code:    fiber.StatusBadRequest,
			}
		}
		return nil, &T.ServiceError{
			Message: "Unable to create user",
			Error:   err,
			Code:    fiber.StatusBadRequest,
		}
	}
	userData, err := R.GetUserByID(userId)
	if err != nil {
		return nil, &T.ServiceError{
			Message: "Unable to create users",
			Error:   err,
			Code:    fiber.StatusBadRequest,
		}
	}

	return &userData, nil
}

func Login(loginDTO *D.LoginDTO) (string, *T.ServiceError) {
	user, err := R.GetUserByEmail(loginDTO.Email)
	if err != nil {
		return "", &T.ServiceError{
			Message: "Unable to create users",
			Error:   err,
			Code:    fiber.StatusUnauthorized,
		}
	}

	email := loginDTO.Email
	pass := loginDTO.Password
	if !ValidatePassword(pass, user.Masterpasswordsalt, user.Masterpasswordhash) {
		return "", &T.ServiceError{
			Message: "Invalid Credentials",
			Error:   err,
			Code:    fiber.StatusUnauthorized,
		}
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(C.Conf.Secret))
	if err != nil {
		return "", &T.ServiceError{
			Message: "Unable to login",
			Error:   err,
			Code:    fiber.StatusUnauthorized,
		}
	}

	return t, nil
}
