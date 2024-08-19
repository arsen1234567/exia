package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	handlers "tender/internal/handlers/oil"
	repositories "tender/internal/repositories/dmart"
	repositoriesprod "tender/internal/repositories/prod"
	dmartService "tender/internal/services/dmart"
	rawService "tender/internal/services/rawData"
)

type application struct {
	errorLog                          *log.Logger
	infoLog                           *log.Logger
	investment_oil_production_handler *handlers.OilReviewHandler
	kgd_taxes_prod_handler            *handlers.OilReviewHandler
	investment_reserves_handler       *handlers.OilReviewHandler
	// permissionHandler  *handlers.PermissionHandler
	// companyHandler     *handlers.CompanyHandler
	// transactionHandler *handlers.TransactionHandler
	// expenseHandler     *handlers.PersonalExpenseHandler
}

func initializeApp(db *sql.DB, errorLog, infoLog *log.Logger) *application {

	investment_oil_production_repository := &repositories.InvestmentOilProductionRepository{Db: db}
	investment_oil_production_service := &dmartService.InvestmentOilProductionService{Repo: investment_oil_production_repository}
	investment_oil_production_handler := &handlers.OilReviewHandler{InvestmentOilProductionService: investment_oil_production_service}

	investment_reserves_repository := &repositories.InvestmentReservesRepository{Db: db}
	investment_reserves_service := &dmartService.InvestmentReservesService{Repo: investment_reserves_repository}
	investment_reserves_handler := &handlers.OilReviewHandler{InvestmentReservesService: investment_reserves_service}

	kgd_taxes_prod_repository := &repositoriesprod.KgdTaxesProdRepository{Db: db}
	kgd_taxes_prod_service := &rawService.KgdTaxesService{Repo: kgd_taxes_prod_repository}
	kgd_taxes_prod_handler := &handlers.OilReviewHandler{KgdTaxesProdService: kgd_taxes_prod_service}

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
		kgd_taxes_prod_handler:            kgd_taxes_prod_handler,
		investment_reserves_handler:       investment_reserves_handler,
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
