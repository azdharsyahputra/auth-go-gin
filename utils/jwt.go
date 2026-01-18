package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte("super-secret-key")

func GenerateJWT(UserID int, role string) (string, error) {
	claims := JWTClaims{
		UserID: UserID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func JwtSecret() []byte {
	return jwtSecret
}
