package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"tender/internal/caching"

	gasHandlers "tender/internal/handlers/gas"
	oilHandlers "tender/internal/handlers/oil"
	dmartRepositories "tender/internal/repositories/dmart"
	prodRepositories "tender/internal/repositories/prod"
	publicRepositories "tender/internal/repositories/public"
	rawDataRepositories "tender/internal/repositories/rawData"
	dmartServices "tender/internal/services/dmart"
	prodServices "tender/internal/services/prod"
	publicServices "tender/internal/services/public"
	rawDataServices "tender/internal/services/rawData"

	newsHandlers "tender/internal/handlers/news"
	reservesHandlers "tender/internal/handlers/reserves"
	newsRepositories "tender/internal/repositories/news"
	newsServices "tender/internal/services/news"
)

type application struct {
	oil_reserves_handler            *oilHandlers.OilReservesHandler
	errorLog                        *log.Logger
	infoLog                         *log.Logger
	oil_review_handler              *oilHandlers.OilReviewHandler
	oil_performance_results_handler *oilHandlers.OilperformanceResultsHandler
	gas_review_handler              *gasHandlers.GasReviewHandler
	oil_benchmarking_handler        *oilHandlers.OilBenchmarkingHandler
	gas_performance_results_handler *gasHandlers.GasperformanceResultsHandler
	media_analysis_handler          *newsHandlers.NewAnalyticsHandler
	reserves_of_oil_handler         *reservesHandlers.ReservesOilHandler
	reserves_of_gas_handler         *reservesHandlers.ReservesGasHandler
	reserves_of_cond_handler        *reservesHandlers.ReservesCondHandler
	// permissionHandler  *handlers.PermissionHandler
	// companyHandler     *handlers.CompanyHandler
	// transactionHandler *handlers.TransactionHandler
	// expenseHandler     *handlers.PersonalExpenseHandler
}

func initializeApp(db *sql.DB, errorLog, infoLog *log.Logger) *application {

	cache := caching.NewInMemoryCache(10*time.Minute, 30*time.Minute)

	reserves_of_cond_repository := &publicRepositories.ReservesCondNgsRepository{Db: db}
	reserves_of_cond_service := &publicServices.ReservesCondNgsService{Repo: reserves_of_cond_repository}

	reserves_of_gas_repository := &publicRepositories.ReservesGasNgsRepository{Db: db}
	reserves_of_gas_service := &publicServices.ReservesGasNgsService{Repo: reserves_of_gas_repository}

	reserves_of_oil_repository := &publicRepositories.ReservesOilNgsRepository{Db: db}
	reserves_of_oil_service := &publicServices.ReservesOilNgsService{Repo: reserves_of_oil_repository}

	google_news_repository := &newsRepositories.GoogleNewsRepository{Db: db}
	google_news_service := &newsServices.GoogleNewsService{Repo: google_news_repository}

	subsoil_geojson_repository := &publicRepositories.SubsoilGeojsonRepository{Db: db}
	subsoil_geojson_service := &publicServices.SubsoilGeojsonService{Repo: subsoil_geojson_repository}

	tax_burden_repository := &prodRepositories.TaxBurdenRepository{Db: db}
	tax_burden_service := &prodServices.TaxBurdenService{Repo: tax_burden_repository, Cache: cache}

	specific_taxes_repository := &prodRepositories.SpecificTaxesRepository{Db: db}
	specific_taxes_service := &prodServices.SpecificTaxesService{Repo: specific_taxes_repository}

	investment_dash_repository := &dmartRepositories.InvestmentsDashRepository{Db: db}
	investment_dash_service := &dmartServices.InvestmentsDashService{Repo: investment_dash_repository}

	production_gas_repository := &prodRepositories.ProductionGasRepository{Db: db}
	production_gas_service := &prodServices.ProductionGasService{Repo: production_gas_repository}

	kgd_taxes_prod_repository := &prodRepositories.KgdTaxesProdRepository{Db: db}
	kgd_taxes_prod_service := &prodServices.KgdTaxesProdService{Repo: kgd_taxes_prod_repository, Cache: cache}

	gas_steps_repository := &prodRepositories.GasStepsRepository{Db: db}
	gas_steps_service := &prodServices.GasStepsService{Repo: gas_steps_repository}

	gas_totals_repository := &prodRepositories.GasTotalsRepository{Db: db}
	gas_totals_service := &prodServices.GasTotalsService{Repo: gas_totals_repository}

	investment_review_forecast_steps_repository := &dmartRepositories.InvestmentReviewForecastStepsRepository{Db: db}
	investment_review_forecast_steps_service := &dmartServices.InvestmentReviewForecastStepsService{Repo: investment_review_forecast_steps_repository}

	investment_review_forecast_total_repository := &dmartRepositories.InvestmentReviewForecastTotalRepository{Db: db}
	investment_review_forecast_total_service := &dmartServices.InvestmentReviewForecastTotalService{Repo: investment_review_forecast_total_repository}

	kgd_taxes_repository := &rawDataRepositories.KgdTaxesRepository{Db: db}
	kgd_taxes_service := &rawDataServices.KgdTaxesService{Repo: kgd_taxes_repository, Cache: cache}

	invest_potential_main_repository := &dmartRepositories.InvestPotentialMainRepository{Db: db}
	invest_potential_main_service := &dmartServices.InvestPotentialMainService{Repo: invest_potential_main_repository}

	ngs_reserves_gas_repository := &rawDataRepositories.NgsReservesGasRepository{Db: db}
	ngs_reserves_gas_service := &rawDataServices.NgsReservesGasService{Repo: ngs_reserves_gas_repository, Cache: cache}

	st_gas_balance_repository := &rawDataRepositories.StGasBalanceRepository{Db: db}
	st_gas_balance_service := &rawDataServices.StGasBalanceService{Repo: st_gas_balance_repository}

	investment_oil_production_repository := &dmartRepositories.InvestmentOilProductionRepository{Db: db}
	investment_oil_production_service := &dmartServices.InvestmentOilProductionService{Repo: investment_oil_production_repository, Cache: cache}

	investment_reserves_repository := &dmartRepositories.InvestmentReservesRepository{Db: db}
	investment_reserves_service := &dmartServices.InvestmentReservesService{Repo: investment_reserves_repository, Cache: cache}

	dfo_gg_reportes_repository := &dmartRepositories.DfoGgReportesRepository{Db: db}
	dfo_gg_reportes_service := &dmartServices.DfoGgReportesService{Repo: dfo_gg_reportes_repository}

	investment_net_profit_repository := &dmartRepositories.InvestmentNetProfitRepository{Db: db}
	investment_net_profit_service := &dmartServices.InvestmentNetProfitService{Repo: investment_net_profit_repository}

	natural_gas_main_repository := &dmartRepositories.NaturalGasMainRepository{Db: db}
	natural_gas_main_service := &dmartServices.NaturalGasMainService{Repo: natural_gas_main_repository}

	dfo_qazaqgas_repository := &dmartRepositories.DfoQazaqgasRepository{Db: db}
	dfo_qazaqgas_service := &dmartServices.DfoQazaqgasService{Repo: dfo_qazaqgas_repository}

	gas_review_handler := &gasHandlers.GasReviewHandler{
		ProductionGasService:  production_gas_service,
		KgdTaxesService:       kgd_taxes_service,
		NgsReservesGasService: ngs_reserves_gas_service,
		DfoGgReportesService:  dfo_gg_reportes_service,
		NaturalGasMainService: natural_gas_main_service,
		GasStepsService:       gas_steps_service,
		GasTotalsService:      gas_totals_service,
		StGasBalanceService:   st_gas_balance_service,
	}

	media_analysis_handler := &newsHandlers.NewAnalyticsHandler{

		GoogleNewsService: google_news_service,
	}
	gas_performance_results_handler := &gasHandlers.GasperformanceResultsHandler{
		DfoQazaqgasService:   dfo_qazaqgas_service,
		DfoGgReportesService: dfo_gg_reportes_service,
	}

	oil_review_handler := &oilHandlers.OilReviewHandler{
		InvestmentOilProductionService:       investment_oil_production_service,
		InvestmentReservesService:            investment_reserves_service,
		KgdTaxesProdService:                  kgd_taxes_prod_service,
		InvestmentNetProfitService:           investment_net_profit_service,
		InvestPotentialMainService:           invest_potential_main_service,
		InvestmentReviewForecastStepsService: investment_review_forecast_steps_service,
		InvestmentReviewForecastTotalService: investment_review_forecast_total_service,
	}

	oil_performance_results_handler := &oilHandlers.OilperformanceResultsHandler{

		InvestmentsDashService: investment_dash_service,
	}

	oil_benchmarking_handler := &oilHandlers.OilBenchmarkingHandler{
		InvestmentsDashService: investment_dash_service,
		SpecificTaxesService:   specific_taxes_service,
		TaxBurdenService:       tax_burden_service,
		KgdTaxesProdService:    kgd_taxes_prod_service,
	}

	oil_reserves_handler := &oilHandlers.OilReservesHandler{

		SubsoilGeojsonService: subsoil_geojson_service,
	}

	reserves_of_gas_handler := &reservesHandlers.ReservesGasHandler{
		ReservesGasService: *reserves_of_gas_service,
	}

	reserves_of_oil_handler := &reservesHandlers.ReservesOilHandler{
		OilReservesService: *reserves_of_oil_service,
	}

	reserves_of_cond_handler := &reservesHandlers.ReservesCondHandler{
		ReservesCondService: *reserves_of_cond_service,
	}
	// permissionRepo := &repositories.PermissionRepository{Db: db}
	// permissionService := &services.PermissionService{Repo: permissionRepo}
	// permissionHandler := &handlers.PermissionHandler{Service: permissionService}

	// companyRepo := &repositories.CompanyRepository{Db: db}
	// companyService := &services.CompanyService{Repo: companyRepo}
	// companyHandler := &handlers.CompanyHandler{Service: companyService}

	// transactionRepo := &repositories.TransactionRepository{Db: db}
	// transactionService := &services.TransactionService{Repo: transactionRepo}
	// transactionHandler := &handlers.TransactionHandler{Service: transactionService}

	// expenseRepo := &repositories.PersonalExpenseRepository{Db: db}
	// expenseService := &services.PersonalExpenseService{Repo: expenseRepo}
	// expenseHandler := &handlers.PersonalExpenseHandler{Service: expenseService}

	return &application{
		oil_reserves_handler:            oil_reserves_handler,
		errorLog:                        errorLog,
		infoLog:                         infoLog,
		gas_review_handler:              gas_review_handler,
		oil_review_handler:              oil_review_handler,
		oil_performance_results_handler: oil_performance_results_handler,
		oil_benchmarking_handler:        oil_benchmarking_handler,
		media_analysis_handler:          media_analysis_handler,
		reserves_of_oil_handler:         reserves_of_oil_handler,
		reserves_of_gas_handler:         reserves_of_gas_handler,
		reserves_of_cond_handler:        reserves_of_cond_handler,
		// permissionHandler:  permissionHandler,
		// companyHandler:     companyHandler,
		// transactionHandler: transactionHandler,
		// expenseHandler:     expenseHandler,

		gas_performance_results_handler: gas_performance_results_handler,
	}

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("%v", err)
		panic("failed to connect to database")
		return nil, err
	}
	db.SetMaxIdleConns(35)
	if err = db.Ping(); err != nil {
		log.Printf("%v", err)
		panic("failed to ping the database")
		return nil, err
	}
	fmt.Println("successfully connected")

	return db, nil
}

func addSecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		w.Header().Set("Cross-Origin-Resource-Policy", "same-origin")
		next.ServeHTTP(w, r)
	})
}
