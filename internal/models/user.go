package models

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	LastName string  `json:"last_name"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	INN      string  `json:"inn"`
	Balance  float64 `json:"balance"`
	Password string  `json:"password"`
}
