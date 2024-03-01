package dto

type TemperatureResponseDTO struct {
	Celsius    float32 `json:"temp_C"`
	Fahrenheit float32 `json:"temp_F"`
	Kelvin     float32 `json:"temp_K"`
}

type currentDTO struct {
	Celsius    float32 `json:"temp_c"`
	Fahrenheit float32 `json:"temp_f"`
}

type WeatherApiResponseDTO struct {
	Current currentDTO `json:"current"`
}
