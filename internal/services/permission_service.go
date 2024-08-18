package services

import (
	"context"
	"tender/internal/models"
	"tender/internal/repositories"
)

type PermissionService struct {
	Repo *repositories.PermissionRepository
}

// AddPermission adds a new permission for a user.
func (s *PermissionService) AddPermission(ctx context.Context, permission models.Permission) (models.Permission, error) {
	permission, err := s.Repo.AddPermission(ctx, permission)
	if err != nil {
		return models.Permission{}, err
	}

	statusValue := 1
	permission.Status = &statusValue
	return permission, nil
}

// DeletePermission deletes a permission for a user.
func (s *PermissionService) DeletePermission(ctx context.Context, id int) error {
	return s.Repo.DeletePermission(ctx, id)
}

// UpdatePermission updates an existing permission.
func (s *PermissionService) UpdatePermission(ctx context.Context, permission models.Permission) (models.Permission, error) {
	return s.Repo.UpdatePermission(ctx, permission)
}

// GetPermission retrieves a permission by ID.
func (s *PermissionService) GetPermissionsByUserID(ctx context.Context, userID int) ([]models.Permission, error) {
	return s.Repo.GetPermissionsByUserID(ctx, userID)
}
