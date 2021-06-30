package order

import (
	"AltaStore/models"
	"AltaStore/models/orderitems"
	// "AltaStore/models/product"
	"time"
)

type Order struct {
	models.GormModel
	// Products   []product.Product `gorm:"many2many:cart_items;" json:"product"`
	// CustomerID uint              `json:"customer_id"`
	StatusPayment  bool
	TotalAmount    int
	OrderDate      time.Time
	Address         string
	StatusDeliver  string
	PhoneNumber    string
	Payment_Method string
	Courier        string
	OrderItems		 []orderitems.OrderItems
	CustomerID     int	`json:"customer_id"`
}
