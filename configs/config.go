package configs

import (
	"AltaStore/models/cart"
	"AltaStore/models/cartitems"
	"AltaStore/models/categories"
	"AltaStore/models/customer"
	"AltaStore/models/product"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// type config struct {
// 	DB_
// }

func InitDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	schema := os.Getenv("SCHEMA")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username, password, host, schema,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	DB.SetupJoinTable(&cart.Cart{}, "Products", &cartitems.CartItems{})

	AutoMigrate()
}

func AutoMigrate() {
	DB.AutoMigrate(&customer.Customer{}, &categories.Category{}, &product.Product{}, &cart.Cart{}, &cartitems.CartItems{})
}
