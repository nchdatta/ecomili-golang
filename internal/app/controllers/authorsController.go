package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nchdatta/ecomili-golang/internal/app/services"
	"github.com/nchdatta/ecomili-golang/internal/app/validations"
	"github.com/nchdatta/ecomili-golang/internal/helpers"
)

func GetAllAuthors(c *fiber.Ctx) error {
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

	authors, err := services.GetAllAuthors(page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			helpers.NewResponse(false, err.Error(), err.Error()),
		)
	}
	return c.JSON(helpers.NewResponse(true, "All Authors List", authors))
}
func GetAuthorByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}

	author, err := services.GetAuthorByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.JSON(helpers.NewResponse(true, "author Details", author))
}

func CreateAuthor(c *fiber.Ctx) error {
	authorCreate := validations.AuthorCreate{}
	if err := c.BodyParser(&authorCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, "Invalid request payload", nil),
		)
	}

	if err := validations.Validate.Struct(authorCreate); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.NewResponse(false, "Invalid User input", err.Error()),
		)
	}

	createdAuthor, err := services.CreateAuthor(&authorCreate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.Status(fiber.StatusCreated).JSON(
		helpers.NewResponse(true, "Author created successfully", createdAuthor),
	)
}

func UpdateAuthor(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}

	authorUpdate := validations.AuthorUpdate{}
	if err := c.BodyParser(&authorUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, "Invalid request payload", nil),
		)
	}

	if err := validations.Validate.Struct(authorUpdate); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.NewResponse(false, "Invalid Author input", err.Error()),
		)
	}

	author, err := services.UpdateAuthor(id, &authorUpdate)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.JSON(helpers.NewResponse(true, "Author Updated.", author))
}
func DeleteAuthor(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}

	author, err := services.DeleteAuthor(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.JSON(helpers.NewResponse(true, "Author Deleted.", author))
}
