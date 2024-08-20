package models

type InvestmentReviewForecastSteps struct {
	Name_short_ru  string  `json:"name_short_ru"`
	Name_short_en  string  `json:"name_short_en"`
	Name_abbr      string  `json:"name_abbr"`
	Year           int     `json:"Year"`
	Brent_price    float64 `json:"Brent_price"`
	OilProduction  float64 `json:"OilProduction"`
	GrossRevenu    float64 `json:"GrossRevenu"`
	GovShare       float64 `json:"GovShare"`
	OperatingCosts float64 `json:"OperatingCosts"`
	EBITDA         float64 `json:"EBITDA"`
	DDA            float64 `json:"DD&A"`
	CIT            float64 `json:"CIT"`
	NetIncome      float64 `json:"NetIncome"`
	CapEx          float64 `json:"CapEx"`
	ATFCF          float64 `json:"ATFCF"`
	Currency       string  `json:"currency"`
	Unit           string  `json:"unit"`
	Scenario       string  `json:"scenario"`
}

type InvestmentReviewForecastStepsSummary struct {
	NameAbbr      string  `json:"name_abbr"`
	Year          int     `json:"Year"`
	OilProduction float64 `json:"OilProduction"`
}
