package city

type GetCityDTO struct {
	City
}

type GetCityInfoDTO struct {
	City
	Country string `json:"country"`
}

type CreateCityDTO struct {
	City
	Country int `json:"country_id"`
}
