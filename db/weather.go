package db

type ReBack struct {
	Message string  `json:"message"`
	Status  int     `json:"status"`
	City    string  `json:"city"`
	Count   int     `json:"count"`
	Data    Weather `json:"data"`
}

type Weather struct {
	Id        int64
	City      string
	Shidu     string     `json:"shidu"`
	Pm25      float64    `json:"pm25"`
	Pm10      float64    `json:"pm10"`
	Quality   string     `json:"quality"`
	Wendu     string     `json:"wendu"`
	Ganmao    string     `json:"ganmao"`
	Yesterday Yesterday  `json:"yesterday" xorm:"-"`
	Forecast  []Forecast `json:"forecast" xorm:"-"`
}

type Yesterday struct {
	CityId  int64
	Date    string  `json:"date"`
	Sunrise string  `json:"sunrise"`
	High    string  `json:"high"`
	Low     string  `json:"low"`
	Sunset  string  `json:"sunset"`
	Aqi     float64 `json:"aqi"`
	Fx      string  `json:"fx"`
	Fl      string  `json:"fl"`
	Type    string  `json:"type"`
	Notice  string  `json:"notice"`
}

type Forecast struct {
	CityId  int64
	Date    string  `json:"date"`
	Sunrise string  `json:"sunrise"`
	High    string  `json:"high"`
	Low     string  `json:"low"`
	Sunset  string  `json:"sunset"`
	Aqi     float64 `json:"aqi"`
	Fx      string  `json:"fx"`
	Fl      string  `json:"fl"`
	Type    string  `json:"type"`
	Notice  string  `json:"notice"`
}
