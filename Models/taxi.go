package models

type TaxiRequest struct {
	Points    []TaxiPoint  `json:"points"`
	Transport string       `json:"transport"`
	Filters   []FilterType `json:"filters"`
	Output    string       `json:"output"`
	Locale    string       `json:"locale"`
}

type TaxiPoint struct {
	Type string  `json:"type"`
	Lon  float64 `json:"lon"`
	Lat  float64 `json:"lat"`
}
