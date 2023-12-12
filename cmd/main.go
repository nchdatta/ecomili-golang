package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nchdatta/ecomili-golang/config"
	"github.com/nchdatta/ecomili-golang/internal/app/routing"
	"github.com/nchdatta/ecomili-golang/internal/database"
)

// Initializing the database
func init() {
	database.ConnectDB()
}

func main() {
	fiberApp := fiber.New()
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	routing.SetUpRoutes(fiberApp)

	log.Fatal(fiberApp.Listen(fmt.Sprintf("%v", config.App.Port)))
}
