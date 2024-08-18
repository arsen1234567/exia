package services

import (
	"context"
	"tender/internal/models"
	"tender/internal/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.Repo.GetAllUsers(ctx)
}

func (s *UserService) SignUp(ctx context.Context, user models.User) (models.User, error) {
	return s.Repo.SignUp(ctx, user)
}

func (s *UserService) LogIn(ctx context.Context, user models.User) (int, error) {
	return s.Repo.LogIn(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (models.User, error) {
	return s.Repo.GetUserByID(ctx, id)
}

func (s *UserService) UpdateBalance(ctx context.Context, id int, amount float64) error {
	return s.Repo.UpdateBalance(ctx, id, amount)
}

func (s *UserService) GetBalance(ctx context.Context, id int) (float64, error) {
	return s.Repo.GetBalance(ctx, id)
}

func (s *UserService) DeleteUserByID(ctx context.Context, id int) error {
	return s.Repo.DeleteUserByID(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	return s.Repo.UpdateUser(ctx, user)
}
