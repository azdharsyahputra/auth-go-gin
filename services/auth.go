package services

import (
	"auth-api/models"
	"auth-api/utils"
	"errors"
)

func Login(email, password string) (*models.User, error) {
	var ErrInvalidCredentials = errors.New("invalid email or password")

	if email == "" || password == "" {
		return nil, errors.New("Email and password are required")
	}

	user, err := models.GetUserByEmail(email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := utils.CheckPassword(user.Password, password); err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
