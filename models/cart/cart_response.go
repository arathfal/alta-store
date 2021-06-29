package cart

type ResponseCart struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    []Cart `json:"data"`
}

type ResponseCartSingle struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Data    Cart `json:"data"`
}
