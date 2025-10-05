package models

type FilterType string

const (
	FilterDirtRoad    FilterType = "dirt_road"
	FilterFerry       FilterType = "ferry"
	FilterHighway     FilterType = "highway"
	FilterBanStairway FilterType = "ban_stairway"
	FilterBanCarRoad  FilterType = "ban_car_road"
	FilterTollRoad    FilterType = "toll_road"
)
