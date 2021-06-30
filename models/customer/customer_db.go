package customer

import (
	"AltaStore/models"
	"AltaStore/models/cart"
	"AltaStore/models/order"
)

type Customer struct {
	models.GormModel
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Cart     cart.Cart
	Order    order.Order
}
