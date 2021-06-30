package cart

import (
	"AltaStore/models"
	"AltaStore/models/product"
)

type Cart struct {
	models.GormModel
	Products   []product.Product `gorm:"many2many:cart_items;" json:"product"`
	CustomerID int              `json:"customer_id"`
}
