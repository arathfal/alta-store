package product

import (
	"AltaStore/models"
)

type Product struct {
	models.GormModel
	Name        string `json:"name"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Description string `json:"desciption"`
	CategoryID  int
}
