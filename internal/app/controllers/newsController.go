package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nchdatta/ecomili-golang/internal/app/services"
	"github.com/nchdatta/ecomili-golang/internal/app/validations"
	"github.com/nchdatta/ecomili-golang/internal/helpers"
)

func GetAllNews(c *fiber.Ctx) error {
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
	newsList, err := services.GetAllNews(page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			helpers.NewResponse(false, err.Error(), err.Error()),
		)
	}
	return c.JSON(helpers.NewResponse(true, "All News List", newsList))
}
func GetNewsByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}

	news, err := services.GetNewsByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.JSON(helpers.NewResponse(true, "News Details", news))
}

func CreateNews(c *fiber.Ctx) error {
	newsCreate := validations.NewsCreate{}
	if err := c.BodyParser(&newsCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, "Invalid request payload", nil),
		)
	}

	if err := validations.Validate.Struct(newsCreate); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.NewResponse(false, "Invalid User input", err.Error()),
		)
	}

	createdNews, err := services.CreateNews(&newsCreate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.Status(fiber.StatusCreated).JSON(
		helpers.NewResponse(true, "News created successfully", createdNews),
	)
}

func UpdateNews(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}

	newsUpdate := validations.NewsUpdate{}
	if err := c.BodyParser(&newsUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, "Invalid request payload", nil),
		)
	}

	if err := validations.Validate.Struct(newsUpdate); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.NewResponse(false, "Invalid Category input", err.Error()),
		)
	}

	news, err := services.UpdateNews(id, &newsUpdate)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.JSON(helpers.NewResponse(true, "News Updated.", news))
}
func DeleteNews(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}

	news, err := services.DeleteNews(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.JSON(helpers.NewResponse(true, "News Deleted.", news))
}
