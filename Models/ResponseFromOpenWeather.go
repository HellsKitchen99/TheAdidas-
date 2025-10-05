package models

type ResponseFromOpenWeather struct {
	Cod     string    `json:"cod"`
	Message int       `json:"message"`
	Cnt     int       `json:"cnt"`
	List    []Weather `json:"list"`
	City    CityInfo  `json:"city"`
}

type Weather struct {
	Dt   int64 `json:"dt"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Visibility int     `json:"visibility"`
	Pop        float64 `json:"pop"`
	Rain       struct {
		ThreeH float64 `json:"3h,omitempty"`
	} `json:"rain,omitempty"`
	Snow struct {
		ThreeH float64 `json:"3h,omitempty"`
	} `json:"snow,omitempty"`
	Sys struct {
		Pod string `json:"pod"`
	} `json:"sys"`
	DtTxt string `json:"dt_txt"`
}

type CityInfo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Coord   struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
	Timezone int `json:"timezone"`
	Sunrise  int `json:"sunrise"`
	Sunset   int `json:"sunset"`
}
