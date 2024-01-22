package middlewares

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/nchdatta/ecomili-golang/config"
)

// JWT Middleware
func JWTMiddleware() func(*fiber.Ctx) error {
	config, _ := config.LoadConfig()

	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.JWT.Secret)},
	})
}
