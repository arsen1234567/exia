package repositories

import (
	"context"
	"database/sql"
	"tender/internal/models"
)

// PersonalExpenseRepository handles database operations for personal expenses.
type PersonalExpenseRepository struct {
	Db *sql.DB
}

// CreatePersonalExpense inserts a new expense into the database.
func (r *PersonalExpenseRepository) CreatePersonalExpense(ctx context.Context, expense models.PersonalExpense) (int, error) {
	result, err := r.Db.ExecContext(ctx, "INSERT INTO personal_expenses (amount, reason, description) VALUES (?, ?, ?)", expense.Amount, expense.Reason, expense.Description)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// GetPersonalExpenseByID retrieves an expense by ID from the database.
func (r *PersonalExpenseRepository) GetPersonalExpenseByID(ctx context.Context, id int) (models.PersonalExpense, error) {
	var expense models.PersonalExpense
	err := r.Db.QueryRowContext(ctx, "SELECT id, amount, reason, description, date FROM personal_expenses WHERE id = ?", id).
		Scan(&expense.ID, &expense.Amount, &expense.Reason, &expense.Description, &expense.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			return expense, models.ErrExpenseNotFound
		}
		return expense, err
	}

	return expense, nil
}

// GetAllPersonalExpenses retrieves all expenses from the database.
func (r *PersonalExpenseRepository) GetAllPersonalExpenses(ctx context.Context) ([]models.PersonalExpense, error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT id, amount, reason, description, date FROM personal_expenses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.PersonalExpense
	for rows.Next() {
		var expense models.PersonalExpense
		err := rows.Scan(&expense.ID, &expense.Amount, &expense.Reason, &expense.Description, &expense.Date)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}

// UpdatePersonalExpense updates an existing expense in the database.
func (r *PersonalExpenseRepository) UpdatePersonalExpense(ctx context.Context, expense models.PersonalExpense) (models.PersonalExpense, error) {
	query := "UPDATE personal_expenses SET"
	params := []interface{}{}

	if expense.Amount != 0 {
		query += " amount = ?,"
		params = append(params, expense.Amount)
	}
	if expense.Reason != "" {
		query += " reason = ?,"
		params = append(params, expense.Reason)
	}
	if expense.Description != "" {
		query += " description = ?,"
		params = append(params, expense.Description)
	}

	// Trim the last comma from the query
	query = query[:len(query)-1]
	query += " WHERE id = ?"
	params = append(params, expense.ID)

	_, err := r.Db.ExecContext(ctx, query, params...)
	if err != nil {
		return models.PersonalExpense{}, err
	}

	// Retrieve the updated expense data
	row := r.Db.QueryRowContext(ctx, "SELECT id, amount, reason, description FROM personal_expenses WHERE id = ?", expense.ID)
	var updatedExpense models.PersonalExpense
	err = row.Scan(&updatedExpense.ID, &updatedExpense.Amount, &updatedExpense.Reason, &updatedExpense.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.PersonalExpense{}, models.ErrExpenseNotFound
		}
		return models.PersonalExpense{}, err
	}

	return updatedExpense, nil
}

// DeletePersonalExpense removes an expense from the database by ID.
func (r *PersonalExpenseRepository) DeletePersonalExpense(ctx context.Context, id int) error {
	result, err := r.Db.ExecContext(ctx, "DELETE FROM personal_expenses WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return models.ErrExpenseNotFound
	}

	return nil
}
