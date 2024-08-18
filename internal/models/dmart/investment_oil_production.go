package models

type InvestmentOilProduction struct {
	Id            int     `json:"id"`
	Name_short_ru string  `json:"name_short_ru"`
	Name_short_en string  `json:"name_short_en"`
	Name_abbr     string  `json:"name_abbr"`
	Abd_scope     bool    `json:"abd_scope"`
	Company       string  `json:"company"`
	Year          int     `json:"year"`
	Value         float64 `json:"value"`
	Unit          string  `json:"unit"`
}



type InvestmentOilProductionSummary struct {
	CoverageScope string  `json:"coverage_scope"`
	Unit          string  `json:"unit"`
	TotalValue    float64 `json:"total_value"`
}
