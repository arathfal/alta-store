package customer

import "AltaStore/models"

type Customer struct {
	models.GormModel
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
