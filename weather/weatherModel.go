package weather

type WeatherDesc struct {
	Value string `json:"value"`
}

type CurrentCondition struct {
	TempC          string        `json:"temp_C"`
	FeelsLikeC     string        `json:"FeelsLikeC"`
	PrecipMM       string        `json:"precipMM"`
	Pressure       string        `json:"pressure"`
	Visibility     string        `json:"visibility"`
	WeatherCode    string        `json:"weatherCode"`
	WindDir16Point string        `json:"winddir16Point"`
	WindDirDegree  string        `json:"winddirDegree"`
	WindSpeedKmph  string        `json:"windspeedKmph"`
	WeatherDesc    []WeatherDesc `json:"weatherDesc"`
}

type AreaInfo struct {
	Value string `json:"value"`
}

type NearestArea struct {
	AreaName   []AreaInfo `json:"areaName"`
	Country    []AreaInfo `json:"country"`
	Latitude   string     `json:"latitude"`
	Longitude  string     `json:"longitude"`
	Population string     `json:"population"`
	Region     []AreaInfo `json:"region"`
}

type WeatherData struct {
	CurrentCondition []CurrentCondition `json:"current_condition"`
	NearestArea      []NearestArea      `json:"nearest_area"`
}
