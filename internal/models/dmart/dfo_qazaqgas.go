package models

type DfoQazaqgas struct {
	ID          int     `json:"id"`
	Item        string  `json:"item"`
	Group       string  `json:"group"`
	Type        string  `json:"type"`
	ReportYear  int     `json:"report_year"`
	PeriodEnd   float64 `json:"period_end"`
	PeriodStart float64 `json:"period_start"`
	IdCompany   int     `json:"id_company"`
	Company     string  `json:"company"`
	ItemEng     string  `json:"item_eng"`
	GroupEng    string  `json:"group_eng"`
	TypeEng     string  `json:"type_eng"`
	CompanyEng  string  `json:"company_eng"`
}

type RevenueByServicesSummary struct {
	RevenueFromGasTransportationServices float64 `json:"revenue_from_gas_transportation_services"`
	RevenueFromGasPipelineMaintenance    float64 `json:"revenue_from_gas_pipeline_maintenance"`
	Other                                float64 `json:"other"`
	GasSalesRevenue                      float64 `json:"gas_sales_revenue"`
	ManagerialFee                        float64 `json:"managerial_fee"`
	TotalRevenue                         float64 `json:"total_revenue"`
}

type RevenueByGeographySummary struct {
	China        float64 `json:"china"`
	Kazakhstan   float64 `json:"kazakhstan"`
	CIS          float64 `json:"cis"`
	TotalRevenue float64 `json:"total_revenue"`
}

type CostItemsSummary struct {
	CostOfSoldGas                   float64 `json:"cost_of_sold_gas"`
	TransportCosts                  float64 `json:"transport_costs"`
	SalaryAndRelatedDeductionsCosts float64 `json:"salary_and_related_deduction_costs"`
	OtherExpenses                   float64 `json:"other_expenses"`
	DepreciationAndAmortization     float64 `json:"depreciation_and_amortization"`
	TotalCosts                      float64 `json:"total_costs"`
}
