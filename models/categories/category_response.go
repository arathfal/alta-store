package categories

type ResponseCategory struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    []Category `json:"data"`
}

type ResponseCategorySingle struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Data    Category `json:"data"`
}
