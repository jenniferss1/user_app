package response

type UserReponse struct {
	Id     string  `json:"id"`
	Email  string  `json:"email"`
	Name   string  `json:"name"`
	Age    int8    `json:"age"`
	Weight float64 `json:"weight"`
	Height float64 `json:"height"`
}
