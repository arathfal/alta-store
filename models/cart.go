package model

type Cart struct {
	GormModel
	Products []Product `gorm:"many2many:cart_products`
}