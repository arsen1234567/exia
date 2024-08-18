package services

import (
	"context"
	"tender/internal/models"
	"tender/internal/repositories"
)

// PersonalExpenseService provides operations on personal expenses.
type PersonalExpenseService struct {
	Repo *repositories.PersonalExpenseRepository
}

// CreatePersonalExpense creates a new personal expense.
func (s *PersonalExpenseService) CreatePersonalExpense(ctx context.Context, expense models.PersonalExpense) (models.PersonalExpense, error) {
	id, err := s.Repo.CreatePersonalExpense(ctx, expense)
	if err != nil {
		return models.PersonalExpense{}, err
	}

	expense.ID = id
	return expense, nil
}

// GetPersonalExpenseByID retrieves a personal expense by ID.
func (s *PersonalExpenseService) GetPersonalExpenseByID(ctx context.Context, id int) (models.PersonalExpense, error) {
	return s.Repo.GetPersonalExpenseByID(ctx, id)
}

// GetAllPersonalExpenses retrieves all personal expenses.
func (s *PersonalExpenseService) GetAllPersonalExpenses(ctx context.Context) ([]models.PersonalExpense, error) {
	return s.Repo.GetAllPersonalExpenses(ctx)
}

// UpdatePersonalExpense updates an existing personal expense.
func (s *PersonalExpenseService) UpdatePersonalExpense(ctx context.Context, expense models.PersonalExpense) (models.PersonalExpense, error) {
	return s.Repo.UpdatePersonalExpense(ctx, expense)
}

// DeletePersonalExpense deletes a personal expense by ID.
func (s *PersonalExpenseService) DeletePersonalExpense(ctx context.Context, id int) error {
	return s.Repo.DeletePersonalExpense(ctx, id)
}
