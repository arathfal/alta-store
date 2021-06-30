package product

import (
	"AltaStore/models"
	"AltaStore/models/categories"
)

type Product struct {
	models.GormModel
	Name        string              `json:"name"`
	Stock       int                 `json:"stock"`
	Price       int                 `json:"price"`
	Description string              `json:"description"`
	CategoryID  uint                `json:"category_id"`
	Category    categories.Category `gorm:"foreignkey:CategoryID"`
}
