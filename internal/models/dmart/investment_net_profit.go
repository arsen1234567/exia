package models

type InvestmentNetProfit struct {
	Company   string  `json:"Company"`
	Year      int     `json:"Year"`
	Currency  string  `json:"currency"`
	NetProfit float64 `json:"NetProfit"`
}

type InvestmentNetProfitSummary struct {
	Year     int64   `json:"Year"`
	TotalSum float64 `json:"Total_sum"`
}
