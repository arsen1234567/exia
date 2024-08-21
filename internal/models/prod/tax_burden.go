package models

type TaxBurden struct {
	Summa         float64 `json:"summa"`
	Name_short_ru string  `json:"Компания"`
	Name_short_en string  `json:"Company"`
	Revenue       float64 `json:"revenue"`
	Currency      int     `json:"currency"`
	Year          int     `json:"year"`
	Report_year   int     `json:"report_year"`
	Currencyunit  string  `json:"currencyunit"`
	Value         float64 `json:"value"`
	Id            int     `json:"id"`
}
