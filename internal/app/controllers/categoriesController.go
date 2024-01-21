package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nchdatta/ecomili-golang/internal/app/services"
	"github.com/nchdatta/ecomili-golang/internal/app/validations"
	"github.com/nchdatta/ecomili-golang/internal/helpers"
)

func GetAllCategories(c *fiber.Ctx) error {
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

	categories, err := services.GetAllCategories(page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			helpers.NewResponse(false, err.Error(), err.Error()),
		)
	}
	return c.JSON(helpers.NewResponse(true, "All Categories List", categories))
}
func GetCategoryByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}

	category, err := services.GetCategoryByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.JSON(helpers.NewResponse(true, "Category Details", category))
}

func CreateCategory(c *fiber.Ctx) error {
	categoryCreate := validations.CategoryCreate{}
	if err := c.BodyParser(&categoryCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, "Invalid request payload", nil),
		)
	}

	if err := validations.Validate.Struct(categoryCreate); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.NewResponse(false, "Invalid User input", err.Error()),
		)
	}

	createdCategory, err := services.CreateCategory(&categoryCreate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.Status(fiber.StatusCreated).JSON(
		helpers.NewResponse(true, "Category created successfully", createdCategory),
	)
}

func UpdateCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}

	categoryUpdate := validations.CategoryUpdate{}
	if err := c.BodyParser(&categoryUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, "Invalid request payload", nil),
		)
	}

	if err := validations.Validate.Struct(categoryUpdate); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.NewResponse(false, "Invalid Category input", err.Error()),
		)
	}

	infobite, err := services.UpdateCategory(id, &categoryUpdate)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.JSON(helpers.NewResponse(true, "Infobite Updated.", infobite))
}
func DeleteCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}

	category, err := services.DeleteCategory(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.JSON(helpers.NewResponse(true, "Category Deleted.", category))
}
