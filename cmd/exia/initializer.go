package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"tender/internal/handlers"
	"tender/internal/repositories"
	"tender/internal/services"
)

type application struct {
	errorLog           *log.Logger
	infoLog            *log.Logger
	userHandler        *handlers.UserHandler
	permissionHandler  *handlers.PermissionHandler
	companyHandler     *handlers.CompanyHandler
	transactionHandler *handlers.TransactionHandler
	expenseHandler     *handlers.PersonalExpenseHandler
}

func initializeApp(db *sql.DB, errorLog, infoLog *log.Logger) *application {

	userRepo := &repositories.UserRepository{Db: db}
	userService := &services.UserService{Repo: userRepo}
	userHandler := &handlers.UserHandler{Service: userService}

	permissionRepo := &repositories.PermissionRepository{Db: db}
	permissionService := &services.PermissionService{Repo: permissionRepo}
	permissionHandler := &handlers.PermissionHandler{Service: permissionService}

	companyRepo := &repositories.CompanyRepository{Db: db}
	companyService := &services.CompanyService{Repo: companyRepo}
	companyHandler := &handlers.CompanyHandler{Service: companyService}

	transactionRepo := &repositories.TransactionRepository{Db: db}
	transactionService := &services.TransactionService{Repo: transactionRepo}
	transactionHandler := &handlers.TransactionHandler{Service: transactionService}

	expenseRepo := &repositories.PersonalExpenseRepository{Db: db}
	expenseService := &services.PersonalExpenseService{Repo: expenseRepo}
	expenseHandler := &handlers.PersonalExpenseHandler{Service: expenseService}

	return &application{
		errorLog:           errorLog,
		infoLog:            infoLog,
		userHandler:        userHandler,
		permissionHandler:  permissionHandler,
		companyHandler:     companyHandler,
		transactionHandler: transactionHandler,
		expenseHandler:     expenseHandler,
	}

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
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
