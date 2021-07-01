package cart

import (
	"AltaStore/models"
	"AltaStore/models/cartitems"
	"AltaStore/models/product"
	// "AltaStore/models/product"
)

type Cart struct {
	models.GormModel
	Products   []product.Product `gorm:"many2many:cart_items;" json:"product"`
	CustomerID uint              `json:"customer_id"`
	// Products   []product.Product `gorm:"many2many:cart_items;" json:"product"`
	CartItems []cartitems.CartItems `json:"cart"`
}
