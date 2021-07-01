package categories

import "AltaStore/models"

type ResponseCategory struct {
	models.Response
	Data []Category `json:"data"`
}

type ResponseCategorySingle struct {
	models.Response
	Data Category `json:"data"`
}
