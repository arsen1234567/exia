package repositories

import (
	"context"
	"database/sql"
	models "tender/internal/models/public"
)

type SubsoilGeojsonRepository struct {
	Db *sql.DB
}

func (r *SubsoilGeojsonRepository) GetSubsoilGeojson(ctx context.Context) ([]models.SubsoilGeojson, error) {
	query := `
		SELECT 
		*
		FROM 
			public.subsoil_geojson
		
		`

	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.SubsoilGeojson
	for rows.Next() {
		var summary models.SubsoilGeojson
		if err := rows.Scan(&summary.Deposit, &summary.Company, &summary.Geojson, &summary.Ploshad); err != nil {
			return nil, err
		}
		results = append(results, summary)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
