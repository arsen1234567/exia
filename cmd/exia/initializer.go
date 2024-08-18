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
	"tender/internal/services"
	prodServices "tender/internal/services/prod"
	rawDataServices "tender/internal/services/rawData"
)

type application struct {
	errorLog                          *log.Logger
	infoLog                           *log.Logger
	investment_oil_production_handler *oilHandlers.InvestmentOilProductionHandler
	gas_review_handler                *gasHandlers.GasReviewHandler
	// permissionHandler  *handlers.PermissionHandler
	// companyHandler     *handlers.CompanyHandler
	// transactionHandler *handlers.TransactionHandler
	// expenseHandler     *handlers.PersonalExpenseHandler
}

func initializeApp(db *sql.DB, errorLog, infoLog *log.Logger) *application {

	investment_oil_production_repository := &dmartRepositories.InvestmentOilProductionRepository{Db: db}
	investment_oil_production_service := &services.InvestmentOilProductionService{Repo: investment_oil_production_repository}
	investment_oil_production_handler := &oilHandlers.InvestmentOilProductionHandler{Service: investment_oil_production_service}
	production_gas_repository := &prodRepositories.ProductionGasRepository{Db: db}
	production_gas_service := &prodServices.ProductionGasService{Repo: production_gas_repository}
	kgd_taxes_repository := &rawDataRepositories.KgdTaxesRepository{Db: db}
	kgd_taxes_service := &rawDataServices.KgdTaxesService{Repo: kgd_taxes_repository}
	gas_review_handler := &gasHandlers.GasReviewHandler{Service: production_gas_service, KgdTaxesService: kgd_taxes_service}

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
		errorLog:                          errorLog,
		infoLog:                           infoLog,
		investment_oil_production_handler: investment_oil_production_handler,
		gas_review_handler:                gas_review_handler,
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
