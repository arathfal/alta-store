package cartproducts

import "AltaStore/models"

type CartProducts struct {
	models.GormModel
	CartID     int
	customerID int
	Quantity   int
}