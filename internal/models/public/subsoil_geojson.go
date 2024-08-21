package models

import "database/sql"

type SubsoilGeojson struct {
	Geojson sql.NullString  `json:"geojson"`
	Deposit sql.NullString  `json:"Deposit"`
	Company sql.NullString  `json:"Компания"`
	Ploshad sql.NullFloat64 `json:"Ploshad"`
}
