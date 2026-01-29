package weather

type Coordinates struct {
	Latitude  float32 `json:"latitude`
	Longitude float32 `json:"longitude`
}

type PointResponse struct {
	Properties Grid `json:"properties"`
}

type Grid struct {
	GridX int `json:"gridX"`
	GridY int `json:"gridY"`
}

type ForecastResponse struct {
	Properties Properties `json:"properties"`
}

type Properties struct {
	Periods []Period `json:"periods"`
}

type Period struct {
	Number        int    `json:"number"`
	Temperature   int    `json:"temperature"`
	ShortForecast string `json:"shortForecast"`
}

type WeatherResponse struct {
	ShortForecast          string `json:"shortForecast"`
	TemperatureDescription string `json:"temperatureDescription"`
}
