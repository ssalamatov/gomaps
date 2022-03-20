package city

import "time"

type City struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	IsCapital  bool      `json:"is_capital"`
	Population int64     `json:"population"`
	FoundAt    time.Time `json:"found_at"`
}
