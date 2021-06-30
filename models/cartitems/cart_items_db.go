package cartitems

import "AltaStore/models/product"

type CartItems struct {
	CartID    int `gorm:"primaryKey"`
	ProductID int `gorm:"primaryKey"`
	Product   product.Product 
	Quantity  int `json:"quantity"`
}
