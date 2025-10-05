package models

type CarRequest struct {
	Points    []PointForCarRequest `json:"points"`
	Transport string               `json:"transport"`
	Filters   []FilterType         `json:"filters"`
	Output    string               `json:"output"`
	Locale    string               `json:"locale"`
}

type PointForCarRequest struct {
	Type string  `json:"type"`
	Lon  float64 `json:"lon"`
	Lat  float64 `json:"lat"`
}
