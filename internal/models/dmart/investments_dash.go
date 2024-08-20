package models

type InvestmentsDash struct {
	name_short_ru               string  `json:"name_short_ru"`
	name_abbr                   string  `json:"name_abbr "`
	Production                  float64 `json:"Production"`
	ProductionUnit              string  `json:"ProductionUnit"`
	report_type                 string  `json:"report_type"`
	report_year                 int     `json:"report_year"`
	Revenue                     int     `json:"Revenue"`
	NetProfit                   int     `json:"NetProfit"`
	Capital                     int     `json:"Capital"`
	ShortAssets                 int     `json:"ShortAssets"`
	ShortLiabilities            int     `json:"ShortLiabilities"`
	FinReserves                 int     `json:"FinReserves"`
	CashAndEquivalents          int     `json:"CashAndEquivalents"`
	FinExpenses                 int     `json:"FinExpenses"`
	EBT                         int     `json:"EBT"`
	LongLoans                   int     `json:"LongLoans"`
	ShortLoans                  int     `json:"ShortLoans"`
	CostOfSales                 int     `json:"CostOfSales"`
	ShortAssetsSale             int     `json:"ShortAssetsSale"`
	LongAssets                  int     `json:"LongAssets"`
	LongLiabilities             int     `json:"LongLiabilities"`
	OperationsCF                int     `json:"OperationsCF"`
	CapExOS                     int     `json:"CapExOS"`
	CapExNMA                    int     `json:"CapExNMA"`
	CashEndPeriod               int     `json:"CashEndPeriod"`
	OperatingProfit             int     `json:"OperatingProfit"`
	DepreciationAndAmortization float64 `json:"DepreciationAndAmortization"`
	TotalTaxes                  float64 `json:"TotalTaxes"`
	USD                         float64 `json:"USD"`
	BalanceАВС1                 float64 `json:"Balance(А+В+С1)"`
	BalanceС2                   float64 `json:"Balance(С2)"`
	GKZАВС1                     float64 `json:"GKZ(А+В+С1)"`
	GKZС2                       float64 `json:"GKZ(С2)"`
	currencyunit                float64 `json:"currencyunit"`
	name_short_en               float64 `json:"name_short_en"`
}
