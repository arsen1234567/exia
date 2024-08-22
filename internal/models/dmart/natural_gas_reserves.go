package models

import "time"

type NaturalGasReserves struct {
	TotalReserv float64   `json:"total_reserves"`
	Year        time.Time `json:"year"`
	Company     string    `json:"Компания"`
	AbdScope    bool      `json:"abd_scope"`
}
