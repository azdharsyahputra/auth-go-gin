package models

import (
	"database/sql"
	"errors"

	"auth-api/config"
)

type User struct {
	ID       int
	Email    string
	Password string
	Role     string
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	row := config.DB.QueryRow(
		"SELECT id, email, password, role FROM users WHERE email = ? LIMIT 1;",
		email,
	)

	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	return &user, nil
}
