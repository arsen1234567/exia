package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	gasHandlers "tender/internal/handlers/gas"
	oilHandlers "tender/internal/handlers/oil"
	dmartRepositories "tender/internal/repositories/dmart"
	prodRepositories "tender/internal/repositories/prod"
	rawDataRepositories "tender/internal/repositories/rawData"
	dmartServices "tender/internal/services/dmart"
	prodServices "tender/internal/services/prod"
	rawDataServices "tender/internal/services/rawData"
)

type application struct {
	errorLog           *log.Logger
	infoLog            *log.Logger
	oil_review_handler *oilHandlers.OilReviewHandler
	gas_review_handler *gasHandlers.GasReviewHandler
	// permissionHandler  *handlers.PermissionHandler
	// companyHandler     *handlers.CompanyHandler
	// transactionHandler *handlers.TransactionHandler
	// expenseHandler     *handlers.PersonalExpenseHandler
}

func initializeApp(db *sql.DB, errorLog, infoLog *log.Logger) *application {

	production_gas_repository := &prodRepositories.ProductionGasRepository{Db: db}
	production_gas_service := &prodServices.ProductionGasService{Repo: production_gas_repository}

	kgd_taxes_prod_repository := &prodRepositories.KgdTaxesProdRepository{Db: db}
	kgd_taxes_prod_service := &prodServices.KgdTaxesProdService{Repo: kgd_taxes_prod_repository}

	gas_steps_repository := &prodRepositories.GasStepsRepository{Db: db}
	gas_steps_service := &prodServices.GasStepsService{Repo: gas_steps_repository}

	gas_totals_repository := &prodRepositories.GasTotalsRepository{Db: db}
	gas_totals_service := &prodServices.GasTotalsService{Repo: gas_totals_repository}

	kgd_taxes_repository := &rawDataRepositories.KgdTaxesRepository{Db: db}
	kgd_taxes_service := &rawDataServices.KgdTaxesService{Repo: kgd_taxes_repository}

	ngs_reserves_gas_repository := &rawDataRepositories.NgsReservesGasRepository{Db: db}
	ngs_reserves_gas_service := &rawDataServices.NgsReservesGasService{Repo: ngs_reserves_gas_repository}

	st_gas_balance_repository := &rawDataRepositories.StGasBalanceRepository{Db: db}
	st_gas_balance_service := &rawDataServices.StGasBalanceService{Repo: st_gas_balance_repository}

	investment_oil_production_repository := &dmartRepositories.InvestmentOilProductionRepository{Db: db}
	investment_oil_production_service := &dmartServices.InvestmentOilProductionService{Repo: investment_oil_production_repository}

	investment_reserves_repository := &dmartRepositories.InvestmentReservesRepository{Db: db}
	investment_reserves_service := &dmartServices.InvestmentReservesService{Repo: investment_reserves_repository}

	dfo_gg_reportes_repository := &dmartRepositories.DfoGgReportesRepository{Db: db}
	dfo_gg_reportes_service := &dmartServices.DfoGgReportesService{Repo: dfo_gg_reportes_repository}

	natural_gas_main_repository := &dmartRepositories.NaturalGasMainRepository{Db: db}
	natural_gas_main_service := &dmartServices.NaturalGasMainService{Repo: natural_gas_main_repository}

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

	oil_review_handler := &oilHandlers.OilReviewHandler{
		InvestmentOilProductionService: investment_oil_production_service,
		InvestmentReservesService:      investment_reserves_service,
		KgdTaxesProdService:            kgd_taxes_prod_service,
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
		errorLog:           errorLog,
		infoLog:            infoLog,
		gas_review_handler: gas_review_handler,
		oil_review_handler: oil_review_handler,
		// permissionHandler:  permissionHandler,
		// companyHandler:     companyHandler,
		// transactionHandler: transactionHandler,
		// expenseHandler:     expenseHandler,

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
