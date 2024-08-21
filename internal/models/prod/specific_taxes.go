package models

type SpecificTaxes struct {
	Name_short_ru  string  `json:"name_short_ru"`
	Name_short_en  string  `json:"name_short_en"`
	Name_abbr      string  `json:"name_abbr"`
	Year_temp      string  `json:"year_temp"`
	Year           int     `json:"year"`
	Production     float64 `json:"Production"`
	ProductionUnit string  `json:"ProductionUnit"`
	Summa          float64 `json:"summa"`
	Specific_tax   float64 `json:"specific_tax"`
}
