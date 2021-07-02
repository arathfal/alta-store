package cart

import "AltaStore/models"

type ResponseCart struct {
	models.Response
	Data []Cart `json:"data"`
}

type ResponseCartSingle struct {
	models.Response
	Data Cart `json:"data"`
}
