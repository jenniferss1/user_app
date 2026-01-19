package request

type UserRequest struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Name     string  `json:"name"`
	Age      int8    `json:"age"`
	Weight   float64 `json:"weight"`
	Height   float64 `json:"height"`
}
