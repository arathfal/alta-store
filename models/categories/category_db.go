package categories

import "AltaStore/models"


type Category struct {
	models.GormModel
	Name    string `json:"name"`
	Product []product.produc
}
