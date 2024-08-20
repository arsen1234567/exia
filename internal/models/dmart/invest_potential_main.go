package models

type InvestPotentialMain struct {
	SpecificOperatingExpenses float64 `json:"Удельные операционные расходы"`
	MarginEBITDA              float64 `json:"Маржа EBITDA"`
	NPV                       float64 `json:"NPV"`
	Taxes                     float64 `json:"Налоги"`
	CapEx                     float64 `json:"CapEx"`
	NPV_production            float64 `json:"NPV_production"`
	InventoryRatio            int     `json:"Кратность запасов"`
}
