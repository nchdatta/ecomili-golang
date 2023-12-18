package database

import (
	"fmt"
	"log"

	"github.com/nchdatta/ecomili-golang/config"
	"github.com/nchdatta/ecomili-golang/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

// Function to connect to the database
func ConnectDB() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed to get database config.\n", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to the database.\n", err)
	}

	log.Println("DB connected!")

	if err := db.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Category{},
		&models.CategoryTag{},
		&models.Infobite{},
		&models.News{},
		&models.Tag{},
		&models.NewsImage{},
	); err != nil {
		log.Fatalln("Error Occured on migration.\n", err)
	}

	log.Println("DB migrated!")
	DBConn = db
}
