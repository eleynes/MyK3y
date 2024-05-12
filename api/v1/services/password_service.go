package services

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"

	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"

	D "github.com/eleynes/MyK3y/data/dto"
	M "github.com/eleynes/MyK3y/models"
	T "github.com/eleynes/MyK3y/types"
)

func GeneratePassword(ctx *fiber.Ctx, generatePasswordDTO *D.GeneratePasswordDTO) (*[]M.Password, *T.ServiceError) {
	var passwords []M.Password
	var err error
	key := []byte("16byteAESKey1234")

	for i := 0; i < generatePasswordDTO.Count; i++ {
		var passwordStr string
		if generatePasswordDTO.PasswordType == "random" {
			passwordStr, err = GenerateSecureRandomPassword(generatePasswordDTO.Length, generatePasswordDTO.IsNumbersIncluded, generatePasswordDTO.IsSymbolsIncluded, generatePasswordDTO.IsUppercaseIncluded)
		} else if generatePasswordDTO.PasswordType == "pin" {
			passwordStr, err = GenerateSecurePIN(generatePasswordDTO.Length)
		} else if generatePasswordDTO.PasswordType == "alphanumeric" {
			passwordStr, err = GenerateSecureAlphanumericPassword(generatePasswordDTO.Length, generatePasswordDTO.IsNumbersIncluded, generatePasswordDTO.IsUppercaseIncluded)
		} else {
			return nil, &T.ServiceError{
				Message: "Invalid password type. Please choose 'random','alphanumeric' or 'pin'.",
				Error:   err,
				Code:    fiber.StatusBadRequest,
			}
		}

		if err != nil {
			return nil, &T.ServiceError{
				Message: "Error generating password/PIN.",
				Error:   err,
				Code:    fiber.StatusBadRequest,
			}
		}

		message := []byte(passwordStr)

		encryptedpassword := Encrypt(message, key)

		password := M.Password{
			ID:                uuid.New(),
			Plainpassword:     passwordStr,
			Encryptedpassword: hex.EncodeToString(encryptedpassword),
		}

		passwords = append(passwords, password)
	}

	return &passwords, nil
}

func generateRandomNumber(max int64) (int64, error) {
	num, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0, err
	}
	return num.Int64(), nil
}

// Function to generate a secure random alphanumeric password
func GenerateSecureRandomPassword(length int, numbers, symbols, uppercase bool) (string, error) {
	var passwordSet string
	if numbers {
		passwordSet += "0123456789"
	}
	if symbols {
		passwordSet += "!@#$%^&*()"
	}
	if uppercase {
		passwordSet += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	passwordSet += "abcdefghijklmnopqrstuvwxyz"

	password := make([]byte, length)
	for i := range password {
		num, err := generateRandomNumber(int64(len(passwordSet)))
		if err != nil {
			return "", err
		}
		password[i] = passwordSet[num]
	}

	return string(password), nil
}

// Function to generate a secure random alphanumeric password
func GenerateSecureAlphanumericPassword(length int, numbers, uppercase bool) (string, error) {
	var passwordSet string
	if numbers {
		passwordSet += "0123456789"
	}
	if uppercase {
		passwordSet += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	passwordSet += "abcdefghijklmnopqrstuvwxyz"

	password := make([]byte, length)
	for i := range password {
		num, err := generateRandomNumber(int64(len(passwordSet)))
		if err != nil {
			return "", err
		}
		password[i] = passwordSet[num]
	}

	return string(password), nil
}

// Function to generate a secure random PIN
func GenerateSecurePIN(length int) (string, error) {
	pin := make([]byte, length)
	for i := range pin {
		num, err := generateRandomNumber(10)
		if err != nil {
			return "", err
		}
		pin[i] = byte(num + 48) // ASCII digits start from 48
	}

	return string(pin), nil
}
