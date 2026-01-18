package services

import (
	"auth-api/models"
	"auth-api/utils"
	"errors"
)

func Login(email, password string) (string, error) {
	var ErrInvalidCredentials = errors.New("invalid email or password")

	if email == "" || password == "" {
		return "", errors.New("Email and password are required")
	}

	user, err := models.GetUserByEmail(email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if err := utils.CheckPassword(user.Password, password); err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
