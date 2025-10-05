package models

type BusRequest struct {
	Locale    string             `json:"locale"`
	Source    Source             `json:"source"`
	Target    Target             `json:"target"`
	Transport []TransportBusType `json:"transport"`
}

type Source struct {
	Name  string             `json:"name"`
	Point PointForBusRequest `json:"point"`
}

type Target struct {
	Name  string             `json:"name"`
	Point PointForBusRequest `json:"point"`
}

type PointForBusRequest struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type TransportBusType string

const (
	Bus  TransportBusType = "bus"
	Tram TransportBusType = "tram"
)
