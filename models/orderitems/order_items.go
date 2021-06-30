package orderitems

type OrderItems struct {
	OrderID 		int
	// ProductID int `gorm:"primaryKey"`
	Quantity    int `json:"quantity"`
	Name        string
	Price       int
	Description int
}
