package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"tender/internal/models"
)

type PermissionRepository struct {
	Db *sql.DB
}

// AddPermission inserts a new permission into the database.
func (r *PermissionRepository) AddPermission(ctx context.Context, permission models.Permission) (models.Permission, error) {
	result, err := r.Db.ExecContext(ctx, "INSERT INTO permissions (user_id, company_id, status) VALUES (?, ?, 1)", permission.UserID, permission.CompanyID)
	if err != nil {
		return models.Permission{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Permission{}, err
	}

	row := r.Db.QueryRowContext(ctx,
		`SELECT permissions.id, user_id, company_id, status, c.name, u.name
				FROM permissions 
				    JOIN companies c on permissions.company_id = c.id 
				    JOIN users u on u.id = permissions.user_id 
				WHERE id = ?`, int(id))

	var createdPermission models.Permission
	err = row.Scan(&createdPermission.ID, &createdPermission.UserID, &createdPermission.CompanyID, &createdPermission.Status, &createdPermission.CompanyName, &createdPermission.UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Permission{}, models.ErrPermissionNotFound
		}
		return models.Permission{}, err
	}

	return createdPermission, nil
}

// DeletePermission removes a permission from the database by ID.
func (r *PermissionRepository) DeletePermission(ctx context.Context, id int) error {
	result, err := r.Db.ExecContext(ctx, "DELETE FROM permissions WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return models.ErrPermissionNotFound
	}

	return nil
}

// UpdatePermission updates an existing permission in the database.
func (r *PermissionRepository) UpdatePermission(ctx context.Context, permission models.Permission) (models.Permission, error) {
	query := "UPDATE permissions SET"
	params := []interface{}{}

	if permission.UserID != 0 {
		query += " user_id = ?,"
		params = append(params, permission.UserID)
	}
	if permission.CompanyID != 0 {
		query += " company_id = ?,"
		params = append(params, permission.CompanyID)
	}
	if permission.Status != nil { // Check if Status is provided in the JSON
		query += " status = ?,"
		params = append(params, *permission.Status)
	}

	// Check if any fields were updated (i.e., if there are any params)
	if len(params) == 0 {
		return models.Permission{}, fmt.Errorf("no fields to update")
	}

	// Trim the last comma from the query string
	query = query[:len(query)-1]
	query += " WHERE id = ?"
	params = append(params, permission.ID)

	_, err := r.Db.ExecContext(ctx, query, params...)
	if err != nil {
		return models.Permission{}, err
	}

	// Retrieve the updated permission data
	row := r.Db.QueryRowContext(ctx,
		`SELECT permissions.id, user_id, company_id, status, c.name, u.name
				FROM permissions 
				    JOIN companies c on permissions.company_id = c.id 
				    JOIN users u on u.id = permissions.user_id 
				WHERE id = ?`, permission.ID)

	var updatedPermission models.Permission
	err = row.Scan(&updatedPermission.ID, &updatedPermission.UserID, &updatedPermission.CompanyID, &updatedPermission.Status, &updatedPermission.CompanyName, &updatedPermission.UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Permission{}, models.ErrPermissionNotFound
		}
		return models.Permission{}, err
	}

	return updatedPermission, nil
}

// GetPermission retrieves a permission by ID from the database.
func (r *PermissionRepository) GetPermissionsByUserID(ctx context.Context, userID int) ([]models.Permission, error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT p.id, p.user_id, p.company_id, p.status, u.name, c.name FROM permissions p JOIN tender.users u on u.id = p.user_id JOIN tender.companies c on c.id = p.company_id WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []models.Permission

	for rows.Next() {
		var permission models.Permission
		err := rows.Scan(&permission.ID, &permission.UserID, &permission.CompanyID, &permission.Status, &permission.UserName, &permission.CompanyName)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}
