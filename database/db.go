package database

import (
	"Zadanie4/models"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DB2 *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("file:database/products.db?_pragma=foreign_keys(ON)"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	if err := DB.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}

	DB2, err = gorm.Open(sqlite.Open("file:database/carts.db?_pragma=foreign_keys(ON)"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	if err := DB2.AutoMigrate(&models.Cart{}, &models.CartItem{}); err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}

}
