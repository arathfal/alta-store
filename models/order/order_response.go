package order

import (
	"AltaStore/models"
	// "AltaStore/models/product"
)

type Checkout struct {
	TotalAmount		int			`json:"total"`
	BankAccount		string	`json:"bank"`
	BankNumber		string	`json:"account_number"`
}

type CheckoutResponse struct {
	models.Response
	// Products   []product.Product `gorm:"many2many:cart_items;" json:"product"`
	// CustomerID uint              `json:"customer_id"`
	Data Checkout `json:"data"`
}
