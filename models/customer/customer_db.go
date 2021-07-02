package customer

import (
	"AltaStore/models"
	"AltaStore/models/cart"
	"AltaStore/models/order"

	"gorm.io/gorm"
)

type Customer struct {
	models.GormModel
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Cart     cart.Cart
	Order    order.Order
}


func (c *Customer) AfterCreate(tx *gorm.DB) (err error) {
	// tx.Model(c).Update(, "admin")
	var cart cart.Cart
	cart.CustomerID = c.ID
	tx.Save(&cart)
  return
}