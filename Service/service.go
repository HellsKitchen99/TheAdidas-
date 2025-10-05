package service

import (
	models "TheAdidasTM/Models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var layout string = "2006-01-02T15:04:05Z"
var apiKey string = "783e0858-39de-4c83-a72c-bc2858c795be"

func EventsProcess(requestData models.RequestData) (models.ResponseToIlya, error) {
	var today []models.Event = requestData.Today
	var tomorrow []models.Event = requestData.Tomorrow
	var responseToIlya models.ResponseToIlya

	var newToday []models.ResponseEventData
	var newTomorrow []models.ResponseEventData

	//итерация по today
	for i := 1; i < len(today); i++ {

		//определение location пред события
		var prevEventLocation string = today[i-1].Location

		//определение location текущего события
		var currEventLocation string = today[i].Location

		//строка запроса к апи (prevLocation)
		var geoCoderUrlPrevLocation string = fmt.Sprintf("https://catalog.api.2gis.com/3.0/items/geocode?q=%v&fields=items.point&key=%v", prevEventLocation, apiKey)

		//строка запроса к апи (curLocation)
		var geoCoderUrlCurrLocation string = fmt.Sprintf("https://catalog.api.2gis.com/3.0/items/geocode?q=%v&fields=items.point&key=%v", currEventLocation, apiKey)

		//запрос к геокодеру (prevLocation)
		responseFromGeoCoderPrevLocation, err := RequestToGeoCoder(geoCoderUrlPrevLocation)
		if err != nil {
			return responseToIlya, err
		}

		//запрос к геокодеру (currLocation)
		responseFromGeoCoderCurrLocation, err := RequestToGeoCoder(geoCoderUrlCurrLocation)
		if err != nil {
			return responseToIlya, err
		}

		//получение координат prevLocation
		var pointsPrevLocation models.Point = responseFromGeoCoderPrevLocation.Result.Items[0].Point
		var prevLat float64 = pointsPrevLocation.Lat
		var prevLon float64 = pointsPrevLocation.Lon

		//получение координат currLocation
		var pointsCurrLocation models.Point = responseFromGeoCoderCurrLocation.Result.Items[0].Point
		var currLat float64 = pointsCurrLocation.Lat
		var currLon float64 = pointsCurrLocation.Lon

		//берем конец пред события и начало текущего
		var endOfPrevEvent string = today[i-1].EndEvent
		var startOfCurrEvent string = today[i].StartEvent

		//парсим время
		endOfPrevEventTime, err := time.Parse(layout, endOfPrevEvent)
		if err != nil {
			return responseToIlya, err
		}
		startOfCurrEventTime, err := time.Parse(layout, startOfCurrEvent)
		if err != nil {
			return responseToIlya, err
		}

		//берем секунды
		endOfPrevEventTimeSeconds := endOfPrevEventTime.Unix()
		startOfCurrEventTimeSeconds := startOfCurrEventTime.Unix()

		//считаем разницу
		diff := startOfCurrEventTimeSeconds - endOfPrevEventTimeSeconds

		transportsTimeToReach, err := MakeRequestForEveryTransport(prevLat, prevLat, currLat, currLon)
		if err != nil {
			return responseToIlya, err
		}

		//запрашиваем погоду
		//var openWeatherUrl string = fmt.Sprintf("api.openweathermap.org/data/2.5/forecast?lat={%v}&lon={%v}&appid={%v}", currLat, currLon, apiKey)
		/*responseFromOpenweather, err := RequestToOpenWeather(openWeatherUrl)
		if err != nil {
			return responseToIlya, err
		}*/

		//weather для Илюхи
		/*temp :=
		feelsLike :=
		condition :=
		windSpeed :=
		windDir :=
		var weather models.WeatherForIlya = models.WeatherForIlya{

		}*/

		var responseEvent models.ResponseEventData = models.ResponseEventData{
			Name: today[i].Name,
			UserLocation: models.Coordinates{
				Lat: prevLat,
				Lon: prevLon,
			},
			EventLocation: models.Coordinates{
				Lat: currLat,
				Lon: currLon,
			},
			ToEventDuration: int(diff),
			TransportTypes:  transportsTimeToReach,
		}
		newToday = append(newToday, responseEvent)
	}

	for i := 1; i < len(tomorrow); i++ {

		//определение location пред события
		prevEventLocation := tomorrow[i-1].Location

		//определение location нынешнего события
		currEventLocation := tomorrow[i].Location

		//строка запроса к апи (prevLocation)
		var geoCoderUrlPrevLocation string = fmt.Sprintf("https://catalog.api.2gis.com/3.0/items/geocode?q=%v&fields=items.point&key=%v", prevEventLocation, apiKey)

		//строка запроса к апи (currLocation)
		var geoCoderUrlCurrLocation string = fmt.Sprintf("https://catalog.api.2gis.com/3.0/items/geocode?q=%v&fields=items.point&key=%v", currEventLocation, apiKey)

		//запрос к геокодеру (prevLocation)
		responseFromGeoCoderPrevLocation, err := RequestToGeoCoder(geoCoderUrlPrevLocation)
		if err != nil {
			return responseToIlya, err
		}

		//запрос к геокодеру (currLocation)
		responseFromGeoCoderCurrLocation, err := RequestToGeoCoder(geoCoderUrlCurrLocation)
		if err != nil {
			return responseToIlya, err
		}

		//получение координат prevLocation
		var pointsPrevLocation models.Point = responseFromGeoCoderPrevLocation.Result.Items[0].Point
		var prevLat float64 = pointsPrevLocation.Lat
		var prevLon float64 = pointsPrevLocation.Lon

		//получение координат currLocation
		var pointsCurrLocation models.Point = responseFromGeoCoderCurrLocation.Result.Items[0].Point
		var currLat float64 = pointsCurrLocation.Lat
		var currLon float64 = pointsCurrLocation.Lon

		//берем конец пред события и начало текущего
		var endOfPrevEvent string = today[i-1].EndEvent
		var startOfCurrEvent string = today[i].StartEvent

		//парсим время
		endOfPrevEventTime, err := time.Parse(layout, endOfPrevEvent)
		if err != nil {
			return responseToIlya, err
		}
		startOfCurrEventTime, err := time.Parse(layout, startOfCurrEvent)
		if err != nil {
			return responseToIlya, err
		}

		//берем секунды
		endOfPrevEventTimeSeconds := endOfPrevEventTime.Unix()
		startOfCurrEventTimeSeconds := startOfCurrEventTime.Unix()

		//считаем разницу
		diff := startOfCurrEventTimeSeconds - endOfPrevEventTimeSeconds

		transportsTimeToReach, err := MakeRequestForEveryTransport(prevLat, prevLat, currLat, currLon)
		if err != nil {
			return responseToIlya, err
		}

		var responseEvent models.ResponseEventData = models.ResponseEventData{
			Name: today[i].Name,
			UserLocation: models.Coordinates{
				Lat: prevLat,
				Lon: prevLon,
			},
			EventLocation: models.Coordinates{
				Lat: currLat,
				Lon: currLon,
			},
			ToEventDuration: int(diff),
			TransportTypes:  transportsTimeToReach,
		}
		newTomorrow = append(newTomorrow, responseEvent)
	}

	responseToIlya = models.ResponseToIlya{
		Today:    newToday,
		Tommorow: newTomorrow,
	}
	return responseToIlya, nil
}

func RequestToGeoCoder(url string) (models.ResponseFromGeoCoder, error) {
	var responseFromGeoCoder models.ResponseFromGeoCoder
	resp, err := http.Get(url)
	if err != nil {
		return responseFromGeoCoder, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return responseFromGeoCoder, err
	}
	if err = json.Unmarshal(data, &responseFromGeoCoder); err != nil {
		return responseFromGeoCoder, err
	}
	return responseFromGeoCoder, nil
}

func MakeRequestForEveryTransport(fromLat, fromLon, toLat, toLon float64) ([]models.TransportForResponseEventData, error) {
	var results []models.TransportForResponseEventData
	transports := []string{"bus", "car", "walking", "scooter", "taxi"}
	for _, mode := range transports {
		switch mode {
		case "bus":
			var apiUrl = fmt.Sprintf("https://routing.api.2gis.com/public_transport/2.0?key=%v", apiKey)
			form := models.BusRequest{
				Locale: "ru",
				Source: models.Source{
					Name: "A",
					Point: models.PointForBusRequest{
						Lat: fromLat,
						Lon: fromLon,
					},
				},
				Target: models.Target{
					Name: "B",
					Point: models.PointForBusRequest{
						Lat: toLat,
						Lon: toLon,
					},
				},
				Transport: []models.TransportBusType{"bus"},
			}
			body, err := json.Marshal(form)
			if err != nil {
				return results, err
			}
			totalDuration, err := RequestToTransportApi(apiUrl, body)
			if err != nil {
				return results, err
			}
			results = append(results, models.TransportForResponseEventData{
				Type:        mode,
				Duration:    totalDuration,
				StatusColor: getColor(totalDuration),
			})
		case "car":
			var apiUrl string = fmt.Sprintf("https://routing.api.2gis.com/routing/7.0.0/global?key=%v", apiKey)
			var point1 models.PointForCarRequest = models.PointForCarRequest{Lat: fromLat, Lon: fromLon, Type: "stop"}
			var point2 models.PointForCarRequest = models.PointForCarRequest{Lat: toLat, Lon: toLon, Type: "stop"}
			var points []models.PointForCarRequest = []models.PointForCarRequest{point1, point2}
			var form models.CarRequest = models.CarRequest{
				Points:    points,
				Transport: "driving",
				Filters: []models.FilterType{
					models.FilterDirtRoad,
					models.FilterTollRoad,
					models.FilterFerry,
				},
				Output: "detailed",
				Locale: "ru",
			}
			body, err := json.Marshal(form)
			if err != nil {
				return results, err
			}
			totalDuration, err := RequestToTransportApi(apiUrl, body)
			if err != nil {
				return results, err
			}
			results = append(results, models.TransportForResponseEventData{
				Type:        mode,
				Duration:    totalDuration,
				StatusColor: getColor(totalDuration),
			})
		case "walking":
			var apiUrl string = fmt.Sprintf("https://routing.api.2gis.com/routing/7.0.0/global?key=%v", apiKey)
			var point1 models.WalkingPoint = models.WalkingPoint{Lat: fromLat, Lon: fromLon, Type: "stop"}
			var point2 models.WalkingPoint = models.WalkingPoint{Lat: toLat, Lon: toLon, Type: "stop"}
			var points []models.WalkingPoint = []models.WalkingPoint{point1, point2}
			var form models.WalkingRequest = models.WalkingRequest{
				Points:    points,
				Transport: "walking",
				Params: models.WalkingParams{
					Pedestrian: models.PedestrianParams{
						UseInstructions: true,
					},
				},
				Filters: []models.FilterType{
					models.FilterDirtRoad,
					models.FilterFerry,
					models.FilterHighway,
					models.FilterBanStairway,
				},
				Output:        "detailed",
				Locale:        "ru",
				NeedAltitudes: true,
			}
			body, err := json.Marshal(form)
			if err != nil {
				return results, err
			}
			totalDuration, err := RequestToTransportApi(apiUrl, body)
			if err != nil {
				return results, err
			}
			results = append(results, models.TransportForResponseEventData{
				Type:        mode,
				Duration:    totalDuration,
				StatusColor: getColor(totalDuration),
			})
		case "scooter":
			var apiUrl string = fmt.Sprintf("https://routing.api.2gis.com/routing/7.0.0/global?key=%v", apiKey)
			var point1 models.ScooterPoint = models.ScooterPoint{Lat: fromLat, Lon: fromLon, Type: "stop"}
			var point2 models.ScooterPoint = models.ScooterPoint{Lat: toLat, Lon: toLon, Type: "stop"}
			var form models.ScooterRequest = models.ScooterRequest{
				Points:    []models.ScooterPoint{point1, point2},
				Transport: "scooter",
				Filters: []models.FilterType{
					models.FilterBanCarRoad,
					models.FilterBanStairway,
				},
				Output:        "detailed",
				Locale:        "ru",
				NeedAltitudes: true,
			}
			body, err := json.Marshal(form)
			if err != nil {
				return results, err
			}
			totalDuration, err := RequestToTransportApi(apiUrl, body)
			if err != nil {
				return results, err
			}
			results = append(results, models.TransportForResponseEventData{
				Type:        mode,
				Duration:    totalDuration,
				StatusColor: getColor(totalDuration),
			})
		case "taxi":
			var apiUrl string = fmt.Sprintf("https://routing.api.2gis.com/routing/7.0.0/global?key=%v", apiKey)
			var point1 models.TaxiPoint = models.TaxiPoint{Lat: fromLat, Lon: fromLon, Type: "stop"}
			var point2 models.TaxiPoint = models.TaxiPoint{Lat: toLat, Lon: toLon, Type: "stop"}
			var form models.TaxiRequest = models.TaxiRequest{
				Points:    []models.TaxiPoint{point1, point2},
				Transport: "taxi",
				Filters: []models.FilterType{
					models.FilterDirtRoad,
					models.FilterTollRoad,
					models.FilterFerry,
				},
				Output: "detailed",
				Locale: "ru",
			}
			body, err := json.Marshal(form)
			if err != nil {
				return results, err
			}
			totalDuration, err := RequestToTransportApi(apiUrl, body)
			if err != nil {
				return results, err
			}
			results = append(results, models.TransportForResponseEventData{
				Type:        mode,
				Duration:    totalDuration,
				StatusColor: getColor(totalDuration),
			})
		}
	}
	return results, nil
}

type DurationResponse struct {
	Result []struct {
		TotalDuration int `json:"total_duration"`
	} `json:"result"`
}

func RequestToTransportApi(url string, form []byte) (int, error) {
	var totalDuration int
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(form))
	if err != nil {
		return totalDuration, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return totalDuration, err
	}
	var durationResponse DurationResponse
	if err := json.Unmarshal(data, &durationResponse); err != nil {
		return totalDuration, err
	}
	return durationResponse.Result[0].TotalDuration, nil
}

func getColor(duration int) string {
	switch {
	case duration == 0:
		return "gray"
	case duration < 1800:
		return "green"
	case duration < 3600:
		return "yellow"
	default:
		return "red"
	}
}

func RequestToOpenWeather(url string) (models.ResponseFromOpenWeather, error) {
	var response models.ResponseFromOpenWeather
	resp, err := http.Get(url)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return response, err
	}
	return response, nil
}
