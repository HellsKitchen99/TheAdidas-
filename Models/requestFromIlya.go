package models

type RequestData struct {
	Today    []Event `json:"today"`
	Tomorrow []Event `json:"tomorrow"`
}

type Event struct {
	Name       string `json:"name"`
	Location   string `json:"location"`
	StartEvent string `json:"start_event"`
	EndEvent   string `json:"end_event"`
}
