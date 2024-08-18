package models

type Permission struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	CompanyID   int     `json:"company_id"`
	Status      *int    `json:"status"`
	UserName    *string `json:"username,omitempty"`
	CompanyName *string `json:"companyname,omitempty"`
}
