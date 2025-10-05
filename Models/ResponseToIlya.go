package models

type ResponseToIlya struct {
	Today    []ResponseEventData `json:"today"`
	Tommorow []ResponseEventData `json:"tomorrow"`
}

type ResponseEventData struct {
	Name            string                          `json:"name"`
	UserLocation    Coordinates                     `json:"user_location"`
	EventLocation   Coordinates                     `json:"event_location"`
	ToEventDuration int                             `json:"to_event_duration"`
	TransportTypes  []TransportForResponseEventData `json:"transport_types"`
	WeatherForIlya  WeatherForIlya                  `json:"weather_for_ilya"`
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type TransportForResponseEventData struct {
	Type        string `json:"type"`
	Duration    int    `json:"duration"`
	StatusColor string `json:"status_color"`
}

type WeatherForIlya struct {
	Temp      string `json:"temp"`
	FeelsLike int    `json:"feels_like"`
	Condition string `json:"condition"`
	WindSpeed int    `json:"wind_speed"`
	WindDir   string `json:"wind_dir"`
}
