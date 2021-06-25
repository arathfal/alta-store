package product

import "AltaStore/models"

type Product struct {
	models.GormModel
	Name        string
	Stock       int
	Description string
}
