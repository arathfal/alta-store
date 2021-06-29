package cartitems

type CartItems struct {
	CartID     int	`gorm:"primaryKey"`
	ProductID 	int `gorm:"primaryKey"`
	Quantity   int	`json:"quantity"`
}