package customer

type ResponseCustomer struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    []Customer `json:"data"`
}

type ResponseCustomerSingle struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Data    Customer `json:"data"`
}
