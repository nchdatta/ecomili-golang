package routing

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nchdatta/ecomili-golang/internal/app/controllers"
)

// Declering all the routes
func SetUpRoutes(app *fiber.App) {
	api := app.Group("api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Role Routes
	role := api.Group("roles")
	role.Get("/", controllers.AllRoles)
	role.Get("/:id", controllers.GetRoleByID)
	role.Post("/create", controllers.CreateRole)
	role.Put("/update/:id", controllers.UpdateRole)
	role.Delete("/delete/:id", controllers.DeleteRole)

	// User Routes
	user := api.Group("users")
	user.Get("/", controllers.AllUsers)
	user.Get("/:id", controllers.GetUserByID)
	user.Post("/create", controllers.CreateUser)
	user.Put("/update/:id", controllers.UpdateUser)
	user.Delete("/delete/:id", controllers.DeleteUser)
}
