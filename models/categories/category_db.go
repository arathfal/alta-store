package categories

import (
	"AltaStore/models"	
	"AltaStore/models/product"
)


type Category struct {
	models.GormModel
	Name string `json:"name"`
	Product []product.Product
}
