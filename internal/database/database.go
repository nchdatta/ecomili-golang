package database

import (
	"fmt"
	"log"

	"github.com/nchdatta/ecomili-golang/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

// Function to connect to the database
func ConnectDB() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed to get database port!\n", err)
	}

	// Formatting Data Source Name (DSN) dynamically
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to the database.\n", err)
	}

	log.Println("DB connected!")

	// if err := db.AutoMigrate(&models.Note{}); err != nil{
	// 	log.Fatalln("Error Occured on migrating.\n", err)
	// }

	log.Println("DB migrated!")
	DBConn = db
}
