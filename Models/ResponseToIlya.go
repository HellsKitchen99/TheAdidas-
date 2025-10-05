package models

type ResponseToIlya struct {
	Today    []ResponseEventData `json:"today"`
	Tommorow []ResponseEventData `json:"tomorrow"`
}

type ResponseEventData struct {
	Name            string         `json:"name"`
	UserLocation    Coordinates    `json:"user_location"`
	EventLocation   Coordinates    `json:"event_location"`
	StartTime       string         `json:"start_time"`
	EndTime         string         `json:"end_time"`
	ToEventDuration int            `json:"to_event_duration"`
	Weather         WeatherForIlya `json:"weather"`
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
	Temp      int    `json:"temp"`
	FeelsLike int    `json:"feels_like"`
	Condition string `json:"condition"`
	WindSpeed int    `json:"wind_speed"`
	WindDir   string `json:"wind_dir"`
}
