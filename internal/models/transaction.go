package models

import "time"

type Transaction struct {
	ID           int       `json:"id"`
	Type         string    `json:"type"`
	TenderNumber *string   `json:"tender_number,omitempty"`
	UserID       int       `json:"user_id"`
	CompanyID    *int      `json:"company_id,omitempty"`
	Organization *string   `json:"organization,omitempty"`
	Amount       float64   `json:"amount"`
	Total        float64   `json:"total"`
	Date         time.Time `json:"date"`
	Status       int       `json:"status"`
	Expenses     []Expense `json:"expenses"`
	UserName     *string   `json:"username,omitempty"`
	CompanyName  *string   `json:"companyname,omitempty"`
}

type Expense struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Amount        float64   `json:"amount"`
	TransactionID int       `json:"transaction_id"`
	Date          time.Time `json:"date"`
}
