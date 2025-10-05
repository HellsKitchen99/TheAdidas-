package service

import (
	models "TheAdidasTM/Models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var layout string = "2006-01-02T15:04:05Z"
var openWeatherLayout = "2006-01-02 15:04:05"
var apiKey string = "783e0858-39de-4c83-a72c-bc2858c795be"
var weatherKey string = "4533e6577b4d8703b24b39df0d1afd5e"

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
		var endOfCurrEvent string = today[i].EndEvent

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

		//запрашиваем погоду
		var openWeatherUrl string = fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%v&lon=%v&appid=%v&units=metric", currLat, currLon, weatherKey)
		responseFromOpenweather, err := RequestToOpenWeather(openWeatherUrl)
		if err != nil {
			return responseToIlya, err
		}

		var chosenWeather models.Weather
		var minDiff int64

		for i, w := range responseFromOpenweather.List {
			var diff int64 = startOfCurrEventTimeSeconds - w.Dt
			if diff < 0 {
				diff = -diff
			}
			if i == 0 {
				minDiff = diff
				chosenWeather = w
			}
			if diff < minDiff {
				minDiff = diff
				chosenWeather = w
			}
		}

		cond := ""
		if len(chosenWeather.Weather) > 0 {
			cond = chosenWeather.Weather[0].Main
		}

		//weather для Илюхи
		var weatherRes models.WeatherForIlya = models.WeatherForIlya{
			Temp:      int(chosenWeather.Main.Temp),
			FeelsLike: int(chosenWeather.Main.FeelsLike),
			Condition: cond,
			WindSpeed: int(chosenWeather.Wind.Speed),
			WindDir:   getWindDirection(chosenWeather.Wind.Deg),
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
			StartTime:       startOfCurrEvent,
			EndTime:         endOfCurrEvent,
			Weather:         weatherRes,
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
		var endOfPrevEvent string = tomorrow[i-1].EndEvent
		var startOfCurrEvent string = tomorrow[i].StartEvent
		var endOfCurrEvent string = tomorrow[i].EndEvent

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

		//запрашиваем погоду
		var openWeatherUrl string = fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%v&lon=%v&appid=%v&units=metric", currLat, currLon, weatherKey)
		responseFromOpenweather, err := RequestToOpenWeather(openWeatherUrl)
		if err != nil {
			return responseToIlya, err
		}

		var chosenWeather models.Weather
		var minDiff int64

		for i, w := range responseFromOpenweather.List {
			var diff int64 = startOfCurrEventTimeSeconds - w.Dt
			if diff < 0 {
				diff = -diff
			}
			if i == 0 {
				minDiff = diff
				chosenWeather = w
			}
			if diff < minDiff {
				minDiff = diff
				chosenWeather = w
			}
		}

		cond := ""
		if len(chosenWeather.Weather) > 0 {
			cond = chosenWeather.Weather[0].Main
		}

		//weather для Илюхи
		var weatherRes models.WeatherForIlya = models.WeatherForIlya{
			Temp:      int(chosenWeather.Main.Temp),
			FeelsLike: int(chosenWeather.Main.FeelsLike),
			Condition: cond,
			WindSpeed: int(chosenWeather.Wind.Speed),
			WindDir:   getWindDirection(chosenWeather.Wind.Deg),
		}

		var responseEvent models.ResponseEventData = models.ResponseEventData{
			Name: tomorrow[i].Name,
			UserLocation: models.Coordinates{
				Lat: prevLat,
				Lon: prevLon,
			},
			EventLocation: models.Coordinates{
				Lat: currLat,
				Lon: currLon,
			},
			ToEventDuration: int(diff),
			StartTime:       startOfCurrEvent,
			EndTime:         endOfCurrEvent,
			Weather:         weatherRes,
		}
		newTomorrow = append(newTomorrow, responseEvent)
	}

	responseToIlya = models.ResponseToIlya{
		Today:    newToday,
		Tommorow: newTomorrow,
	}
	return responseToIlya, nil
}

func getWindDirection(degrees int) string {
	switch {
	case degrees >= 338 || degrees < 23:
		return "N"
	case degrees >= 23 && degrees < 68:
		return "NE"
	case degrees >= 68 && degrees < 113:
		return "E"
	case degrees >= 113 && degrees < 158:
		return "SE"
	case degrees >= 158 && degrees < 203:
		return "S"
	case degrees >= 203 && degrees < 248:
		return "SW"
	case degrees >= 248 && degrees < 293:
		return "W"
	case degrees >= 293 && degrees < 338:
		return "NW"
	default:
		return ""
	}
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

type DurationResponse struct {
	Result []struct {
		TotalDuration int `json:"total_duration"`
	} `json:"result"`
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
