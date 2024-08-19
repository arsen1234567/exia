package models

type KdgTaxesProd struct {
	Bin               string  `json:"bin"`
	Tax_org_code      int     `json:"tax_org_code"`
	Name_tax_ru       string  `json:"name_tax_ru"`
	Name_tax_kz       string  `json:"name_tax_kz"`
	Kbk               int     `json:"kbk"`
	Kbk_name_ru       string  `json:"kbk_name_ru"`
	Kbk_name_kz       string  `json:"kbk_name_kz"`
	Pay_num           string  `json:"pay_num"`
	Pay_type          string  `json:"pay_type"`
	Entry_type        string  `json:"entry_type"`
	Summa             float32 `json:"summa"`
	Pay_Status        string  `json:"pay_status"`
	Company_ru        string  `json:"Компания"`
	Name_abbr         string  `json:"name_abbr"`
	Budget            string  `json:"budget"`
	Budget_detail     string  `json:"budget_detail"`
	Budget_eng        string  `json:"budget_eng"`
	Budget_detail_eng string  `json:"budget_detail_eng"`
	Currency          string  `json:"currency"`
	Company_en        string  `json:"company"`
}

type KgdTaxesProdSummary struct {
	Year       int     `"json:"year"`
	TotalValue float64 `json:"total_value"`
}
