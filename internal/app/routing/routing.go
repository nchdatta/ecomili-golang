package routing

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nchdatta/ecomili-golang/internal/app/controllers"
	"github.com/nchdatta/ecomili-golang/internal/middlewares"
)

// Declering all the routes
func SetUpRoutes(app *fiber.App) {
	api := app.Group("api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Auth Routes
	auth := api.Group("auth")
	auth.Post("/login", controllers.Login)

	// Role Routes
	role := api.Group("roles")
	role.Use(middlewares.JWTMiddleware())
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

	// Infobite Routes
	infobite := api.Group("infobites")
	infobite.Get("/", controllers.GetAllInfobites)
	infobite.Get("/:id", controllers.GetInfobiteByID)
	infobite.Post("/create", controllers.CreateInfobite)
	infobite.Put("/update/:id", controllers.UpdateInfobite)
	infobite.Delete("/delete/:id", controllers.DeleteInfobite)

	// category Routes
	category := api.Group("categories")
	category.Get("/", controllers.GetAllCategories)
	category.Get("/:id", controllers.GetCategoryByID)
	category.Post("/create", controllers.CreateCategory)
	category.Put("/update/:id", controllers.UpdateCategory)
	category.Delete("/delete/:id", controllers.DeleteCategory)

	// News Routes
	news := api.Group("news")
	news.Get("/", controllers.GetAllNews)
	news.Get("/:id", controllers.GetNewsByID)
	news.Post("/create", controllers.CreateNews)
	news.Put("/update/:id", controllers.UpdateNews)
	news.Delete("/delete/:id", controllers.DeleteNews)

	// Author Routes
	author := api.Group("authors")
	author.Get("/", controllers.GetAllAuthors)
	author.Get("/:id", controllers.GetAuthorByID)
	author.Post("/create", controllers.CreateAuthor)
	author.Put("/update/:id", controllers.UpdateAuthor)
	author.Delete("/delete/:id", controllers.DeleteAuthor)

}
