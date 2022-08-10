package YandexWeather

type Response struct {
	Now       int       `json:"now"`
	NowDt     string    `json:"now_dt"`
	Info      *info     `json:"info"`
	Fact      fact      `json:"fact"`
	Forecasts forecasts `json:"forecasts"`
}

type info struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
	Url string  `json:"url"`
}

type fact struct {
	Temp      int    `json:"temp"`
	FeelsLike int    `json:"feels_like"`
	TempWater int    `json:"temp_water"`
	Icon      string `json:"icon"`
	Condition string `json:"condition"`
	Humidity  int    `json:"humidity"`
}

type forecasts struct {
	Date  string  `json:"date"`
	Parts []parts `json:"parts"`
}

type parts struct {
	PartName  string `json:"part_name"`
	TempMin   int    `json:"temp_min"`
	TempMax   int    `json:"temp_max"`
	FeelsLike int    `json:"feels_like"`
	Icon      string `json:"icon"`
	Condition string `json:"condition"`
	Humidity  int    `json:"humidity"`
}
