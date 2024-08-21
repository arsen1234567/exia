package models

type SubsoilGeojson struct {
	Geojson string  `json:"geojson"`
	Deposit string  `json:"Deposit"`
	Company string  `json:"Компания"`
	Ploshad float64 `json:"Ploshad"`
}
