package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	mux := pat.New()

	// USERS
	mux.Post("/users/signup", http.HandlerFunc(app.userHandler.SignUp))                   // sign up
	mux.Post("/users/login", http.HandlerFunc(app.userHandler.LogIn))                     // login
	mux.Get("/users", http.HandlerFunc(app.userHandler.GetAllUsers))                      // get all users
	mux.Get("/users/details/:id", http.HandlerFunc(app.userHandler.GetUserByID))          // get one user info
	mux.Del("/users/:id", http.HandlerFunc(app.userHandler.DeleteUserByID))               // delete user by id
	mux.Put("/users/:id", http.HandlerFunc(app.userHandler.UpdateUser))                   // update user by id
	mux.Get("/users/balance/:id", http.HandlerFunc(app.userHandler.GetBalance))           // get user balance by id
	mux.Put("/users/balance/update/:id", http.HandlerFunc(app.userHandler.UpdateBalance)) // update user balance

	// PERMISSIONS
	mux.Post("/permissions", http.HandlerFunc(app.permissionHandler.AddPermission))                       // add a new permission
	mux.Get("/permissions/user/:user_id", http.HandlerFunc(app.permissionHandler.GetPermissionsByUserID)) // get all permissions by user ID
	mux.Put("/permissions/:id", http.HandlerFunc(app.permissionHandler.UpdatePermission))                 // update a permission by id
	mux.Del("/permissions/:id", http.HandlerFunc(app.permissionHandler.DeletePermission))                 // delete a permission by id

	// COMPANY
	mux.Post("/companies", http.HandlerFunc(app.companyHandler.CreateCompany))     // Create a new company
	mux.Get("/companies", http.HandlerFunc(app.companyHandler.GetAllCompanies))    // Get all companies
	mux.Get("/companies/:id", http.HandlerFunc(app.companyHandler.GetCompanyByID)) // Get company by ID
	mux.Put("/companies/:id", http.HandlerFunc(app.companyHandler.UpdateCompany))  // Update company by ID
	mux.Del("/companies/:id", http.HandlerFunc(app.companyHandler.DeleteCompany))  // Delete company by ID

	// TRANSACTION
	mux.Post("/transactions", http.HandlerFunc(app.transactionHandler.CreateTransaction))             // Create a new transaction
	mux.Get("/transactions", http.HandlerFunc(app.transactionHandler.GetAllTransactions))             // Get all transactions
	mux.Get("/transactions/:id", http.HandlerFunc(app.transactionHandler.GetTransactionByID))         // Get transaction by ID
	mux.Get("/transactions/user/:id", http.HandlerFunc(app.transactionHandler.GetTransactionsByUser)) // Get transaction by user ID
	mux.Put("/transactions/:id", http.HandlerFunc(app.transactionHandler.UpdateTransaction))          // Update transaction by ID
	mux.Del("/transactions/:id", http.HandlerFunc(app.transactionHandler.DeleteTransaction))          // Delete transaction by ID

	// PERSONAL EXPENSES
	mux.Post("/expenses", http.HandlerFunc(app.expenseHandler.CreatePersonalExpense))     // Create a new expense
	mux.Get("/expenses", http.HandlerFunc(app.expenseHandler.GetAllPersonalExpenses))     // Get all expenses
	mux.Get("/expenses/:id", http.HandlerFunc(app.expenseHandler.GetPersonalExpenseByID)) // Get expense by ID
	mux.Put("/expenses/:id", http.HandlerFunc(app.expenseHandler.UpdatePersonalExpense))  // Update expense by ID
	mux.Del("/expenses/:id", http.HandlerFunc(app.expenseHandler.DeletePersonalExpense))  // Delete expense by ID

	// REPORTS
	// company month
	mux.Get("/reports/company/month/global", http.HandlerFunc(app.transactionHandler.GetMonthlyAmountsByGlobal))               //global - company - month
	mux.Get("/reports/company/month/year", http.HandlerFunc(app.transactionHandler.GetMonthlyAmountsByYear))                   //year - company - month
	mux.Get("/reports/company/month/year/company", http.HandlerFunc(app.transactionHandler.GetMonthlyAmountsByYearAndCompany)) //year and company - company - month

	return mux
}
