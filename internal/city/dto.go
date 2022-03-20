package city

type CreateCityDTO struct {
	City
}
type GetCityDTO struct {
	City
}
type GetCityInfoDTO struct {
	City
	Country string `json:"country"`
}
