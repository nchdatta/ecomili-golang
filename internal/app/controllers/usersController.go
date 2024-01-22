package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nchdatta/ecomili-golang/internal/app/services"
	"github.com/nchdatta/ecomili-golang/internal/app/validations"
	"github.com/nchdatta/ecomili-golang/internal/helpers"
)

func AllUsers(c *fiber.Ctx) error {
	// Page & Limit Query
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	pageSize, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}

	users, err := services.GetAllUsers(page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			helpers.NewResponse(false, err.Error(), err.Error()),
		)
	}

	return c.JSON(helpers.NewResponse(true, "All Users List", users))
}
func GetUserByID(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}

	role, err := services.GetUserByID(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.JSON(helpers.NewResponse(true, "User Details", role))
}

func CreateUser(c *fiber.Ctx) error {
	userCreate := validations.UserCreate{}
	if err := c.BodyParser(&userCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, "Invalid request payload", nil),
		)
	}

	if err := validations.Validate.Struct(userCreate); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.NewResponse(false, "Invalid User input", err.Error()),
		)
	}

	createdUser, err := services.CreateUser(&userCreate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.Status(fiber.StatusCreated).JSON(
		helpers.NewResponse(true, "User created successfully", createdUser),
	)
}

func UpdateUser(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}

	userUpdate := validations.UserUpdate{}
	if err := c.BodyParser(&userUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, "Invalid request payload", nil),
		)
	}

	if err := validations.Validate.Struct(userUpdate); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.NewResponse(false, "Invalid User input", err.Error()),
		)
	}

	user, err := services.UpdatedUser(userId, &userUpdate)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.JSON(helpers.NewResponse(true, "User Updated.", user))
}
func DeleteUser(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}

	user, err := services.DeleteUser(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.JSON(helpers.NewResponse(true, "User Deleted.", user))
}
