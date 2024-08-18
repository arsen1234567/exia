package models

import "time"

type PersonalExpense struct {
	ID          int       `json:"id"`
	Amount      float64   `json:"amount"`
	Reason      string    `json:"reason"`
	Description string    `json:"description,omitempty"`
	Date        time.Time `json:"date"`
}
