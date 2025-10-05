package models

type ScooterRequest struct {
	Points        []ScooterPoint `json:"points"`
	Transport     string         `json:"transport"`
	Filters       []FilterType   `json:"filters"`
	Output        string         `json:"output"`
	Locale        string         `json:"locale"`
	NeedAltitudes bool           `json:"need_altitudes"`
}

type ScooterPoint struct {
	Type string  `json:"type"`
	Lon  float64 `json:"lon"`
	Lat  float64 `json:"lat"`
}
