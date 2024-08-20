package models

type StGasBalance struct {
	ItemNumber string  `json:"item_number"`
	ItemGroup  string  `json:"item_group"`
	Item       string  `json:"item"`
	Unit       string  `json:"unit"`
	IdCompany  int64   `json:"id_company"`
	Year       int64   `json:"year"`
	Value      float64 `json:"value"`
	Id         int64   `json:"id"`
	ItemEn     string  `json:"item_en"`
	UnitEng    string  `json:"unit_eng"`
}

type GasBalanceSummary struct {
	Year                     int64   `json:"year"`
	TotalGasProduction       float64 `json:"total_gas_production"`
	Import                   float64 `json:"import"`
	Export                   float64 `json:"export"`
	PersonalNeeds            float64 `json:"personal_needs"`
	ReinjectionIntoReservoir float64 `json:"reinjection_into_reservoir"`
	GasSales                 float64 `json:"gas_sales"`
	Difference               float64 `json:"difference"`
}
