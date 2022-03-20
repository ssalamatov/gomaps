package city

import "time"

type GetCityDTO struct {
	City
}

type GetCityInfoDTO struct {
	City
	Country string `json:"country"`
}

type CreateCityDTO struct {
	Name       string    `json:"name" binding:"required"`
	IsCapital  bool      `json:"is_capital"`
	Population int64     `json:"population" binding:"required"`
	FoundAt    time.Time `json:"found_at" binding:"required"`
	Country    string    `json:"country" binding:"required"`
}
