package models

type ProductionGas struct {
	ID                   int64   `json:"id"`
	Bin                  string  `json:"bin"`
	Year                 int64   `json:"year"`
	Company              string  `json:"company"`
	ProductionAssociated float64 `json:"production_associated"`
	ProductionNatural    float64 `json:"production_natural"`
	UseOwnNeeds          float64 `json:"use_own_needs"`
	ReverseInjection     float64 `json:"reverse_injection"`
	UseOwnNeedsGTU       float64 `json:"use_own_needs_gtu"`
	SupplyUpgGpz         float64 `json:"supply_upg_gpz"`
	GasLosses            float64 `json:"gas_losses"`
	GazSales             float64 `json:"gaz_sales"`
	GasTopipeline        float64 `json:"gas_topipeline"`
	Burned               float64 `json:"burned"`
}

type ProductionGasSummary struct {
	Year            int64   `json:"year"`
	TotalProduction float64 `json:"total_production"`
}
