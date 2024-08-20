package models

type InvestmentReviewForecastTotals struct {
	Name_short_ru string  `json:"name_short_ru"`
	Name_short_en string  `json:"name_short_en"`
	Name_abbr     string  `json:"name_abbr"`
	Scenario      string  `json:"scenario"`
	Unit          string  `json:"unit"`
	NPV           float64 `json:"NPV"`
	Currency      string  `json:"currency"`
}
