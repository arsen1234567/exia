package models

type GasTotals struct {
	PriceCoef          float64 `json:"price_coef"`
	CompanyID          int     `json:"company_id"`
	Username           string  `json:"username"`
	Currency           string  `json:"currency"`
	NPV                float64 `json:"npv"`
	TerminalValue      float64 `json:"terminal_value"`
	NpvPlusTerminalVal float64 `json:"npv_plus_terminal_val"`
}
