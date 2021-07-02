package customer

import "AltaStore/models"

type ResponseCustomer struct {
	models.Response
	Data []Customer `json:"data"`
}

type ResponseCustomerSingle struct {
	models.Response
	Data Customer `json:"data"`
}

type LoginResponse struct {
	models.Response
	Token string `json:"token"`
}
