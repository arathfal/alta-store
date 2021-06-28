package cartproducts

import "AltaStore/models"

type CartProducts struct {
	models.GormModel
	CartID     int		
	CustomerID int	
	Quantity   int	`json:"quantity"`
}