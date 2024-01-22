package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nchdatta/ecomili-golang/internal/app/services"
	"github.com/nchdatta/ecomili-golang/internal/app/validations"
	"github.com/nchdatta/ecomili-golang/internal/helpers"
)

func Login(c *fiber.Ctx) error {
	authLogin := validations.LoginCredentials{}
	if err := c.BodyParser(&authLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, "Invalid request payload", nil),
		)
	}

	// Now validating the fields
	if err := validations.Validate.Struct(authLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}

	token, err := services.AuthLogin(&authLogin)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}

	return c.Status(200).JSON(
		helpers.NewResponse(true, "Token", token),
	)

}
