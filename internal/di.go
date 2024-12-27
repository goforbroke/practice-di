package di

import (
	"github.com/goforbroke/practice-di/infra/weather"
	"github.com/goforbroke/practice-di/usecase"
)

type Container struct {
	cache map[string]any
}

func NewContainer() *Container {
	return &Container{
		cache: map[string]any{},
	}
}

// Usecase
func (c *Container) UsecaseQuery(place string) {
	uf := usecase.NewUsecaseQuery(c.WeatherForecastClient())
	uf.QueryToWeatherForecast(place)
}

// Infra
func (c *Container) WeatherForecastClient() usecase.WeatherForecastClient {
	return weather.NewWeatherForecastClient()
}
