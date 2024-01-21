package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nchdatta/ecomili-golang/internal/app/services"
	"github.com/nchdatta/ecomili-golang/internal/app/validations"
	"github.com/nchdatta/ecomili-golang/internal/helpers"
)

func GetAllInfobites(c *fiber.Ctx) error {
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

	infobites, err := services.GetAllInfobites(page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			helpers.NewResponse(false, err.Error(), err.Error()),
		)
	}
	return c.JSON(helpers.NewResponse(true, "All infobites List", infobites))
}
func GetInfobiteByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}

	role, err := services.GetInfobiteByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.JSON(helpers.NewResponse(true, "Infobite Details", role))
}

func CreateInfobite(c *fiber.Ctx) error {
	infobiteCreate := validations.InfobiteCreate{}
	if err := c.BodyParser(&infobiteCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, "Invalid request payload", nil),
		)
	}

	if err := validations.Validate.Struct(infobiteCreate); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.NewResponse(false, "Invalid User input", err.Error()),
		)
	}

	createdInfote, err := services.CreateInfobite(&infobiteCreate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.Status(fiber.StatusCreated).JSON(
		helpers.NewResponse(true, "Infobite created successfully", createdInfote),
	)
}

func UpdateInfobite(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}

	infobiteUpdate := validations.InfobiteUpdate{}
	if err := c.BodyParser(&infobiteUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, "Invalid request payload", nil),
		)
	}

	if err := validations.Validate.Struct(infobiteUpdate); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.NewResponse(false, "Invalid User input", err.Error()),
		)
	}

	infobite, err := services.UpdateInfobite(id, &infobiteUpdate)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.JSON(helpers.NewResponse(true, "Infobite Updated.", infobite))
}
func DeleteInfobite(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}

	infobite, err := services.DeleteInfobite(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			helpers.NewResponse(false, err.Error(), nil),
		)
	}
	return c.JSON(helpers.NewResponse(true, "Infobite Deleted.", infobite))
}
