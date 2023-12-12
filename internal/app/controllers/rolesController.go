package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nchdatta/ecomili-golang/internal/app/services"
	"github.com/nchdatta/ecomili-golang/internal/helpers"
)

func AllRoles(c *fiber.Ctx) error {
	return c.SendString("All Roles")
}
func GetRoleByID(c *fiber.Ctx) error {
	roleId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(helpers.NewResponse(false, "Role Not Found.", nil))
	}

	role, err := services.GetRoleByID(roleId)
	if err != nil {
		return c.JSON(helpers.NewResponse(false, "Something went wrong.", err))
	}

	return c.JSON(helpers.NewResponse(true, "Got Role.", role))
}
func CreateRole(c *fiber.Ctx) error {
	return c.SendString("Create Role")
}
func UpdateRole(c *fiber.Ctx) error {
	return c.SendString("Update Role")
}
func DeleteRole(c *fiber.Ctx) error {
	return c.SendString("Delete Role")
}
