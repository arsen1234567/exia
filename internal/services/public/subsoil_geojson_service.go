package services

import (
	"context"
	models "tender/internal/models/public"
	repositories "tender/internal/repositories/public"
)

type SubsoilGeojsonService struct {
	Repo *repositories.SubsoilGeojsonRepository
}

func (s *SubsoilGeojsonService) GetSubsoilGeojson(ctx context.Context) ([]models.SubsoilGeojson, error) {
	return s.Repo.GetSubsoilGeojson(ctx)
}
