package product

type ProductCreate struct {
	Name        string `json:"name"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	CategoryID  uint   `json:"category_id"`
}
