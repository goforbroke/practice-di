package usecase

type UsecaseQuery interface {
	QueryToWeatherForecast(place string) Weather
}

type WeatherForecastClient interface {
	Forecast(place string) (*Weather, error)
}

type usecaseQuery struct {
	client WeatherForecastClient
}

type Weather struct {
	PublicTime string
	Summary    string
	Forecasts  map[string]string
}

func NewUsecaseQuery(client WeatherForecastClient) UsecaseQuery {
	return &usecaseQuery{client: client}
}

func (uq *usecaseQuery) QueryToWeatherForecast(place string) Weather {
	weather, _ := uq.client.Forecast(place)

	// fmt.Printf("Weather summary: %v\n", weather.Summary)
	// fmt.Printf("Weather forecasts...\n")
	// for k, v := range weather.Forecasts {
	// 	fmt.Printf("\t%v: %v\n", k, v)
	// }
	// fmt.Printf("Update: %v", weather.PublicTime)
	return *weather
}
