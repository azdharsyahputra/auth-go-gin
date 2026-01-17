package services

import (
	"auth-api/models"
	"auth-api/utils"
	"errors"
)

var ErrEmailAlreadyExists = errors.New("email already exists")

func Register(email, password string) error {
	if email == "" || password == "" {
		return errors.New("Email and password are required")
	}

	if _, err := models.GetUserByEmail(email); err == nil {
		return ErrEmailAlreadyExists
	}

	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	return models.CreateUser(email, hash)
}
