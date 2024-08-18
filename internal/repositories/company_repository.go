package repositories

import (
	"context"
	"database/sql"
	"tender/internal/models"
)

type CompanyRepository struct {
	Db *sql.DB
}

// CreateCompany inserts a new company into the database.
func (r *CompanyRepository) CreateCompany(ctx context.Context, company models.Company) (int, error) {
	result, err := r.Db.ExecContext(ctx, "INSERT INTO companies (name, description) VALUES (?, ?)", company.Name, company.Description)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// DeleteCompany removes a company from the database by ID.
func (r *CompanyRepository) DeleteCompany(ctx context.Context, id int) error {
	result, err := r.Db.ExecContext(ctx, "DELETE FROM companies WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return models.ErrCompanyNotFound
	}

	return nil
}

// UpdateCompany updates an existing company in the database.
func (r *CompanyRepository) UpdateCompany(ctx context.Context, company models.Company) (models.Company, error) {
	query := "UPDATE companies SET"
	params := []interface{}{}

	if company.Name != "" {
		query += " name = ?,"
		params = append(params, company.Name)
	}
	if company.Description != "" {
		query += " description = ?,"
		params = append(params, company.Description)
	}

	// Trim the last comma from the query
	query = query[:len(query)-1]
	query += " WHERE id = ?"
	params = append(params, company.ID)

	_, err := r.Db.ExecContext(ctx, query, params...)
	if err != nil {
		return models.Company{}, err
	}

	// Retrieve the updated company data
	row := r.Db.QueryRowContext(ctx, "SELECT id, name, description FROM companies WHERE id = ?", company.ID)
	var updatedCompany models.Company
	err = row.Scan(&updatedCompany.ID, &updatedCompany.Name, &updatedCompany.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Company{}, models.ErrCompanyNotFound
		}
		return models.Company{}, err
	}

	return updatedCompany, nil
}

// GetCompanyByID retrieves a company by ID from the database.
func (r *CompanyRepository) GetCompanyByID(ctx context.Context, id int) (models.Company, error) {
	var company models.Company
	err := r.Db.QueryRowContext(ctx, "SELECT id, name, description FROM companies WHERE id = ?", id).Scan(&company.ID, &company.Name, &company.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return company, models.ErrCompanyNotFound
		}
		return company, err
	}

	return company, nil
}

// GetAllCompanies retrieves all companies from the database.
func (r *CompanyRepository) GetAllCompanies(ctx context.Context) ([]models.Company, error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT id, name, description FROM companies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []models.Company
	for rows.Next() {
		var company models.Company
		err := rows.Scan(&company.ID, &company.Name, &company.Description)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}
