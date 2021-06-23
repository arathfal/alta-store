package customer

type CustomerLogin struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
