package product

import "AltaStore/models"

type ProductResponse struct {
	models.Response
	Data []Product `json:"data"`
}

type ProductResponseSingle struct {
	models.Response
	Data Product `json:"data"`
}
