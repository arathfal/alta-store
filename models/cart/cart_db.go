package cart

import "AltaStore/models"

type Cart struct {
	models.GormModel
	Products   []Product `gorm:"many2many:cart_products;"`
	CustomerID uint
}