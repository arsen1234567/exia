package services

import (
	"context"
	"tender/internal/models"
	"tender/internal/repositories"
)

type TransactionService struct {
	Repo *repositories.TransactionRepository
}

// CreateTransaction creates a new transaction with expenses.
func (s *TransactionService) CreateTransaction(ctx context.Context, transaction models.Transaction) (models.Transaction, error) {
	return s.Repo.CreateTransaction(ctx, transaction)
}

// GetTransactionByID retrieves a transaction by ID along with its expenses.
func (s *TransactionService) GetTransactionByID(ctx context.Context, id int) (models.Transaction, error) {
	return s.Repo.GetTransactionByID(ctx, id)
}

// GetAllTransactions retrieves all transactions.
func (s *TransactionService) GetAllTransactions(ctx context.Context) ([]models.Transaction, error) {
	return s.Repo.GetAllTransactions(ctx)
}

// GetTransactionsByUser retrieves transactions by user id.
func (s *TransactionService) GetTransactionsByUser(ctx context.Context, userID int) ([]models.Transaction, error) {
	return s.Repo.GetTransactionsByUser(ctx, userID)
}

// UpdateTransaction updates an existing transaction and its expenses.
func (s *TransactionService) UpdateTransaction(ctx context.Context, transaction models.Transaction) (models.Transaction, error) {
	return s.Repo.UpdateTransaction(ctx, transaction)
}

// DeleteTransaction deletes a transaction and its expenses by ID.
func (s *TransactionService) DeleteTransaction(ctx context.Context, id int) error {
	return s.Repo.DeleteTransaction(ctx, id)
}

func (s *TransactionService) GetMonthlyAmountsByGlobal(ctx context.Context) ([]repositories.YearlyAmounts, error) {
	return s.Repo.GetMonthlyAmountsByGlobal(ctx)
}

func (s *TransactionService) GetMonthlyAmountsByYear(ctx context.Context, year int) ([]repositories.MonthlyAmount, error) {
	return s.Repo.GetMonthlyAmountsByYear(ctx, year)
}

func (s *TransactionService) GetMonthlyAmountsByYearAndCompany(ctx context.Context, year int, companyID int) ([]repositories.MonthlyAmount, error) {
	return s.Repo.GetMonthlyAmountsByYearAndCompany(ctx, year, companyID)
}

func (s *TransactionService) GetMonthlyAmountsGroupedByYear(ctx context.Context) ([]repositories.YearlyAmounts, error) {
	return s.Repo.GetMonthlyAmountsGroupedByYear(ctx)
}

func (s *TransactionService) GetMonthlyAmountsGroupedByYearForUser(ctx context.Context, userID int) ([]repositories.YearlyAmounts, error) {
	return s.Repo.GetMonthlyAmountsGroupedByYearForUser(ctx, userID)
}

func (s *TransactionService) GetMonthlyAmountsForUserByYear(ctx context.Context, userID int, year int) ([]repositories.MonthlyAmount, error) {
	return s.Repo.GetMonthlyAmountsForUserByYear(ctx, userID, year)
}

func (s *TransactionService) GetMonthlyAmountsForUserByYearAndCompany(ctx context.Context, userID int, year int, companyID int) ([]repositories.MonthlyAmount, error) {
	return s.Repo.GetMonthlyAmountsForUserByYearAndCompany(ctx, userID, year, companyID)
}

func (s *TransactionService) GetTotalAmountGroupedByCompany(ctx context.Context) ([]repositories.CompanyTotalAmount, error) {
	return s.Repo.GetTotalAmountGroupedByCompany(ctx)
}

func (s *TransactionService) GetTotalAmountByCompanyForYear(ctx context.Context, year int) ([]repositories.CompanyTotalAmount, error) {
	return s.Repo.GetTotalAmountByCompanyForYear(ctx, year)
}

func (s *TransactionService) GetTotalAmountByCompanyForYearAndMonth(ctx context.Context, year int, month int) ([]repositories.CompanyTotalAmount, error) {
	return s.Repo.GetTotalAmountByCompanyForYearAndMonth(ctx, year, month)
}

func (s *TransactionService) GetTotalAmountGroupedByCompanyForUsers(ctx context.Context) ([]repositories.CompanyTotalAmount, error) {
	return s.Repo.GetTotalAmountGroupedByCompany(ctx)
}

func (s *TransactionService) GetTotalAmountByCompanyForUser(ctx context.Context, userID int) ([]repositories.CompanyTotalAmount, error) {
	return s.Repo.GetTotalAmountByCompanyForUser(ctx, userID)
}

func (s *TransactionService) GetTotalAmountByCompanyForUserAndYear(ctx context.Context, userID int, year int) ([]repositories.CompanyTotalAmount, error) {
	return s.Repo.GetTotalAmountByCompanyForUserAndYear(ctx, userID, year)
}

func (s *TransactionService) GetTotalAmountByCompanyForUserYearAndMonth(ctx context.Context, userID int, year int, month int) ([]repositories.CompanyTotalAmount, error) {
	return s.Repo.GetTotalAmountByCompanyForUserYearAndMonth(ctx, userID, year, month)
}
