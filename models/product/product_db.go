package product

import "AltaStore/models"

type Product struct {
	models.GormModel
	Name        string	`json:"name"`
	Stock       int			`json:"stock"`
	Description string	`json:"desciption"`
}
