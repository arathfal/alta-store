package product

import (
	"AltaStore/models"	
	// "AltaStore/models/categories"
)


type Product struct {
	models.GormModel
	Name        string `json:"name"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	// Category categories.Category
	CategoryID 	int			
}
