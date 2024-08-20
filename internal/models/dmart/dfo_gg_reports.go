package models

type DfoGgReportes struct {
	ID             int     `json:"id"`
	Item           string  `json:"item"`
	Code           string  `json:"code"`
	PeriodEnd      float64 `json:"period_end"`
	PeriodStart    float64 `json:"period_start"`
	IdDFO          int     `json:"id_dfo"`
	ReportDate     string  `json:"report_date"`
	ReportYear     int     `json:"report_year"`
	ReportID       int     `json:"report_id"`
	ReportType     string  `json:"report_type"`
	ReportForm     string  `json:"report_form"`
	ReportTemplate int     `json:"report_template"`
	CfMethod       string  `json:"cf_method"`
	DateApproval   string  `json:"date_approval"`
	IdCompany      int     `json:"id_company"`
	Company        string  `json:"company"`
	ItemEn         string  `json:"item_en"`
	CompanyNameRu  string  `json:"company_name_ru"`
}

type NetProfitSummary struct {
	Year      int     `json:"year"`
	NetProfit float64 `json:"net_profit"`
}
