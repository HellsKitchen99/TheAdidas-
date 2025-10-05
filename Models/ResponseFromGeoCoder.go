package models

type ResponseFromGeoCoder struct {
	Meta   Meta   `json:"meta"`
	Result Result `json:"result"`
}

type Meta struct {
	ApiVersion string `json:"api_version"`
	Code       int    `json:"code"`
	IssueDate  string `json:"issue_date"`
}

type Result struct {
	Items []Item `json:"items"`
	Total int    `json:"total"`
}

type Item struct {
	AddressName string `json:"address_name"`
	FullName    string `json:"full_name"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Point       Point  `json:"point"`
	PurposeName string `json:"purpose_name"`
	Type        string `json:"type"`
}

type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
