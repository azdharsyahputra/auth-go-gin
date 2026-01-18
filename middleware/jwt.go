package middleware

import (
	"net/http"
	"strings"

	"auth-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		parts := strings.Fields(authHeader)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Authorization header format",
			})
			return
		}

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing Authorization header",
			})
			return
		}

		tokenString := parts[1]

		token, err := jwt.ParseWithClaims(
			tokenString,
			&utils.JWTClaims{},
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return utils.JwtSecret(), nil
			},
		)

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "invalid token",
			})
			return
		}

		claims, ok := token.Claims.(*utils.JWTClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token claims",
			})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()

	}
}
