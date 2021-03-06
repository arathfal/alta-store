package configs

import (
	"AltaStore/models/cart"
	"AltaStore/models/cartitems"
	"AltaStore/models/categories"
	"AltaStore/models/customer"
	"AltaStore/models/order"
	"AltaStore/models/orderitems"
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

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		username, password, host, schema,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	// DB.SetupJoinTable(&cart.Cart{}, "Products", &cartitems.CartItems{})

	AutoMigrate()
}

func AutoMigrate() {
	DB.AutoMigrate(&customer.Customer{})
	DB.AutoMigrate(&cart.Cart{})
	DB.AutoMigrate(&product.Product{})
	DB.AutoMigrate(&categories.Category{})
	DB.AutoMigrate(&cartitems.CartItems{})
	DB.AutoMigrate(&order.Order{})
	DB.AutoMigrate(&orderitems.OrderItems{})
}

func InitDBTest() {
	username := "hbstudent"
	password := "hbstudent"
	host := "127.0.0.1:3306"
	schema := "testaltastore"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		username, password, host, schema,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	// DB.SetupJoinTable(&cart.Cart{}, "Products", &cartitems.CartItems{})

	// AutoMigrateTest()
}

// func AutoMigrateTest() {
// 	DB.Migrator().DropTable(&orderitems.OrderItems{}, &order.Order{}, &cartitems.CartItems{},
// 		&categories.Category{}, &product.Product{}, &cart.Cart{}, &customer.Customer{})
// 	DB.AutoMigrate(&customer.Customer{})
// 	DB.AutoMigrate(&cart.Cart{})
// 	DB.AutoMigrate(&product.Product{})
// 	DB.AutoMigrate(&categories.Category{})
// 	DB.AutoMigrate(&cartitems.CartItems{})
// 	DB.AutoMigrate(&order.Order{})
// 	DB.AutoMigrate(&orderitems.OrderItems{})

// }
