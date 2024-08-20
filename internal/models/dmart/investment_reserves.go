package models

type InvestmentReserves struct {
	Id            int     `json:"id"`
	Name_short_ru string  `json:"name_short_ru"`
	Name_short_en string  `json:"name_short_en"`
	Name_abbr     string  `json:"name_abbr"`
	Abd_scope     bool    `json:"abd_scope"`
	Nedrouser     string  `json:"Недропользователь"`
	ABC           float64 `json:"Балансовые запасы на конец(А+В+С1)"`
	C2            float64 `json:"Балансовые запасы на конец(С2)"`
	Type          string  `json:"Тип"`
	Unit          string  `json:"unit"`
	Year          int     `json:"year"`
}

type InvestmentReservesSummary struct {
	CoverageScope string   `json:"coverage_scope"`
	TotalValue    *float64 `json:"total_value"`
}
