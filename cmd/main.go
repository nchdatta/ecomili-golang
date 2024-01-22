package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nchdatta/ecomili-golang/config"
	"github.com/nchdatta/ecomili-golang/internal/app/routing"
	"github.com/nchdatta/ecomili-golang/internal/database"
)

// Initializing the DB
func init() {
	database.ConnectDB()
}

func main() {
	app := fiber.New()
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// Initialize default config
	app.Use(recover.New())
	app.Use(cors.New())

	// Setting up routes
	routing.SetUpRoutes(app)

	// Running the server
	log.Fatal(app.Listen(fmt.Sprintf(":%v", config.App.Port)))
	log.Printf("Server Running on port: %v", config.App.Port)
}
