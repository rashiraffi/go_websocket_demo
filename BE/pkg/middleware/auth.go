package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract JWT from query parameters
		tokenStr := c.Query("token") // Extract from ?token=<JWT>

		if tokenStr == "" {
			// If token is not in query params, fallback to Authorization header
			tokenStr = c.Get("Authorization")
			if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
				tokenStr = tokenStr[7:]
			}
		}

		// Validate JWT token
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil // Use the same signing key
		})

		if err != nil || !token.Valid {
			return c.SendStatus(fiber.StatusUnauthorized) // 401 Unauthorized
		}

		// Store user in context for later use
		c.Locals("user", token)
		return c.Next()
	}
}
