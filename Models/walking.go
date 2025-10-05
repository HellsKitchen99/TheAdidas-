package models

type WalkingRequest struct {
	Points        []WalkingPoint `json:"points"`
	Transport     string         `json:"transport"`
	Params        WalkingParams  `json:"params"`
	Filters       []FilterType   `json:"filters"`
	Output        string         `json:"output"`
	Locale        string         `json:"locale"`
	NeedAltitudes bool           `json:"need_altitudes"`
}

type WalkingPoint struct {
	Type string  `json:"type"`
	Lon  float64 `json:"lon"`
	Lat  float64 `json:"lat"`
}

type WalkingParams struct {
	Pedestrian PedestrianParams `json:"pedestrian"`
}

type PedestrianParams struct {
	UseInstructions bool `json:"use_instructions"`
}
