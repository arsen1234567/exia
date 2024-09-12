package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	mux := pat.New()

	mux.Get("/api/oil/review/financial_indicators/production", http.HandlerFunc(app.oil_review_handler.GetInvestmentOilProductionSummary)) // UPDATED
	mux.Get("/api/oil/review/financial_indicators/reserves", http.HandlerFunc(app.oil_review_handler.GetInvestmentReservesSummary))
	mux.Get("/api/oil/review/financial_indicators/taxes", http.HandlerFunc(app.oil_review_handler.GetKgdTaxesProd)) // NEED TO UPD
	mux.Get("/api/oil/review/financial_indicators/investment_profit", http.HandlerFunc(app.oil_review_handler.GetInvestmentNetProfitSummary))

	mux.Get("/api/oil/review/forecast/reserves", http.HandlerFunc(app.oil_review_handler.GetInvestPotentialMainHandler))
	mux.Get("/api/oil/review/forecast/taxes", http.HandlerFunc(app.oil_review_handler.GetInvestmentReviewForecastStepsSummary))
	mux.Get("/api/oil/review/forecast/production_forecast", http.HandlerFunc(app.oil_review_handler.GetCompaniesForecastSteps))
	mux.Get("/api/oil/review/forecast/ebitda", http.HandlerFunc(app.oil_review_handler.GetEbitdaToGrossRevenueRatio))
	mux.Get("/api/oil/review/forecast/npv", http.HandlerFunc(app.oil_review_handler.GetInvestmentReviewForecastTotal))
	mux.Get("/api/oil/review/forecast/specopex", http.HandlerFunc(app.oil_review_handler.GetSpecOpEx))
	mux.Get("/api/oil/review/forecast/revenue", http.HandlerFunc(app.oil_review_handler.GetRevenueForecastSteps))
	mux.Get("/api/oil/review/forecast/capex", http.HandlerFunc(app.oil_review_handler.GetCapExForecastSteps))
	mux.Get("/api/oil/review/forecast/atfcf", http.HandlerFunc(app.oil_review_handler.GetATFCFForecastSteps))
	mux.Get("/api/oil/review/forecast/opex", http.HandlerFunc(app.oil_review_handler.GetOpExForecastSteps))
	mux.Get("/api/oil/review/forecast/gov_share", http.HandlerFunc(app.oil_review_handler.GetGovShareForecastSteps))

	mux.Get("/api/oil/performance/reserves", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDash))
	mux.Get("/api/oil/performance/production", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashOilProduction))
	mux.Get("/api/oil/performance/specific_revenue", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashSpecificRevenue))
	mux.Get("/api/oil/performance/specific_net_profit", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashSpecificNetProfit))
	mux.Get("/api/oil/performance/roa", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashROA))
	mux.Get("/api/oil/performance/net_profit_margin", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashNetProfitMargin))
	mux.Get("/api/oil/performance/revenue", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashRevenue))
	mux.Get("/api/oil/performance/operating_profit", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashOperatingProfit))
	mux.Get("/api/oil/performance/ebitda", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashEBITDA))
	mux.Get("/api/oil/performance/net_profit", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashNetProfit))
	mux.Get("/api/oil/performance/total_taxes", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashTotalTaxes))
	mux.Get("/api/oil/performance/tax_burden", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashTaxBurden))
	mux.Get("/api/oil/performance/specific_taxes", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashSpecificTaxes))
	mux.Get("/api/oil/performance/assets", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashAssets))
	mux.Get("/api/oil/performance/equity", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashCapital))
	mux.Get("/api/oil/performance/liabilities", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashLiabilities))
	//
	mux.Get("/api/oil/performance/current_ratio", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashCurrentRatio))
	mux.Get("/api/oil/performance/cf", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashCF))
	mux.Get("/api/oil/performance/capex", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashCapEx))
	mux.Get("/api/oil/performance/cash", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashCashEndPeriod))

	mux.Get("/api/news/analysis/companies", http.HandlerFunc(app.media_analysis_handler.GetCompanyCount))
	mux.Get("/api/news/analysis/media", http.HandlerFunc(app.media_analysis_handler.GetMediaCount))
	mux.Get("/api/news/analysis/media-news", http.HandlerFunc(app.media_analysis_handler.GetMediaNews))
	mux.Get("/api/news/analysis/sentimental", http.HandlerFunc(app.media_analysis_handler.GetSentimentCounts))
	mux.Get("/api/news/analysis/sentimental_days", http.HandlerFunc(app.media_analysis_handler.GetSentimentDayCounts))
	mux.Get("/api/news/analysis/sentimental_map", http.HandlerFunc(app.media_analysis_handler.GetSentimentMap))
	mux.Get("/api/news/analysis/top_companies", http.HandlerFunc(app.media_analysis_handler.GetTopCompanyCount))
	mux.Get("/api/news/analysis/top_companies_dict", http.HandlerFunc(app.media_analysis_handler.GetTopCompanyCountDict))
	mux.Get("/api/news/analysis/source_count", http.HandlerFunc(app.media_analysis_handler.GetSourceCount))
	mux.Get("/api/news/analysis/news", http.HandlerFunc(app.media_analysis_handler.GetNewsByPeriod))

	mux.Get("/api/oil/benchmarking/specific_net_profit", http.HandlerFunc(app.oil_benchmarking_handler.GetInvestmentsDashSpecificNetProfitGraph)) // WORK
	mux.Get("/api/oil/benchmarking/roa", http.HandlerFunc(app.oil_benchmarking_handler.GetInvestmentsDashROAGraph))                               // WORK
	mux.Get("/api/oil/benchmarking/specific_taxes", http.HandlerFunc(app.oil_benchmarking_handler.GetSpecificTaxesGraph))                         // WORK
	mux.Get("/api/oil/benchmarking/tax_burden_graph", http.HandlerFunc(app.oil_benchmarking_handler.GetTaxBurdenGraph))                           // UPDATING
	mux.Get("/api/oil/benchmarking/sum_taxes", http.HandlerFunc(app.oil_benchmarking_handler.GetSummaAllTaxes))

	// RESERVES START

	mux.Get("/api/reserves/oil/deposits", http.HandlerFunc(app.reserves_of_oil_handler.GetDeposit))
	mux.Get("/api/reserves/oil/companies", http.HandlerFunc(app.reserves_of_oil_handler.GetNumberOfCompanies))
	mux.Get("/api/reserves/oil/production", http.HandlerFunc(app.reserves_of_oil_handler.GetReservesOilNgsTotalProduction))
	mux.Get("/api/reserves/oil/deposit_region", http.HandlerFunc(app.reserves_of_oil_handler.GetNumberOfDepositsByRegion))
	mux.Get("/api/reserves/oil/top_companies", http.HandlerFunc(app.reserves_of_oil_handler.GetTopCompaniesByReserves))

	mux.Get("/api/reserves/gas/deposits", http.HandlerFunc(app.reserves_of_gas_handler.GetDeposit))
	mux.Get("/api/reserves/gas/companies", http.HandlerFunc(app.reserves_of_gas_handler.GetNumberOfCompanies))
	mux.Get("/api/reserves/gas/total-reserves", http.HandlerFunc(app.reserves_of_gas_handler.GetTotalReserves))
	mux.Get("/api/reserves/gas/production", http.HandlerFunc(app.reserves_of_gas_handler.GetProduction))
	mux.Get("/api/reserves/gas/deposit_region", http.HandlerFunc(app.reserves_of_gas_handler.GetNumberOfDepositsByRegion))
	mux.Get("/api/reserves/gas/top_companies", http.HandlerFunc(app.reserves_of_gas_handler.GetTopCompaniesByReserves))

	mux.Get("/api/reserves/cond/deposits", http.HandlerFunc(app.reserves_of_cond_handler.GetDeposit))
	mux.Get("/api/reserves/cond/companies", http.HandlerFunc(app.reserves_of_cond_handler.GetNumberOfCompanies))
	mux.Get("/api/reserves/cond/total-reserves", http.HandlerFunc(app.reserves_of_cond_handler.GetTotalReserves))
	mux.Get("/api/reserves/cond/production", http.HandlerFunc(app.reserves_of_cond_handler.GetProduction))
	mux.Get("/api/reserves/cond/deposit_region", http.HandlerFunc(app.reserves_of_cond_handler.GetNumberOfDepositsByRegion))
	mux.Get("/api/reserves/cond/top_companies", http.HandlerFunc(app.reserves_of_cond_handler.GetTopCompaniesByReserves))

	// RESERVES END

	mux.Get("/api/gas/review/financial_indicators/production", http.HandlerFunc(app.gas_review_handler.GetGasProductionSummary))
	mux.Get("/api/gas/review/financial_indicators/taxes", http.HandlerFunc(app.gas_review_handler.GetKgdTaxesSummary))                                  // UPDATED
	mux.Get("/api/gas/review/financial_indicators/recoverable_gas_reserves", http.HandlerFunc(app.gas_review_handler.GetRecoverableGasReservesSummary)) // NEED TO UPD
	mux.Get("/api/gas/review/financial_indicators/net_profit", http.HandlerFunc(app.gas_review_handler.GetNetProfitSummary))

	mux.Get("/api/gas/review/perspective/reserve_ratio", http.HandlerFunc(app.gas_review_handler.GetReserveRatio))
	mux.Get("/api/gas/review/perspective/taxes", http.HandlerFunc(app.gas_review_handler.GetAmountOfPredictedTaxes))
	mux.Get("/api/gas/review/perspective/NPV+TV", http.HandlerFunc(app.gas_review_handler.GetNPVplusTV))
	mux.Get("/api/gas/review/perspective/EBITDAmargin", http.HandlerFunc(app.gas_review_handler.GetEBITDAmargin))
	mux.Get("/api/gas/review/perspective/gas_balance", http.HandlerFunc(app.gas_review_handler.GetGasBalance))

	mux.Get("/api/gas/performance/service_revenue", http.HandlerFunc(app.gas_performance_results_handler.GetRevenueByServiceAndCompanyAndYear))
	mux.Get("/api/gas/performance/cost_items", http.HandlerFunc(app.gas_performance_results_handler.GetCostItemsByCompanyAndYear))
	mux.Get("/api/gas/performance/CIT", http.HandlerFunc(app.gas_performance_results_handler.GetCIT))
	mux.Get("/api/gas/performance/geography_revenue", http.HandlerFunc(app.gas_performance_results_handler.GetRevenueByGeographyAndCompanyAndYear))
	mux.Get("/api/gas/performance/revenue", http.HandlerFunc(app.gas_performance_results_handler.GetRevenueByCompanyAndYear))
	mux.Get("/api/gas/performance/cost", http.HandlerFunc(app.gas_performance_results_handler.GetCostOfGoodsWorksServicesSold))
	mux.Get("/api/gas/performance/gross_profit", http.HandlerFunc(app.gas_performance_results_handler.GetGrossProfit))

	return mux
}
