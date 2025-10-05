package service

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}

func GetApiKey() (string, error) {
	apiKey := os.Getenv("APIKEY")
	if apiKey == "" {
		return apiKey, fmt.Errorf("api key is nil")
	}
	return apiKey, nil
}

func GetWeatherKey() (string, error) {
	weatherKey := os.Getenv("WEATHERKEY")
	if weatherKey == "" {
		return weatherKey, fmt.Errorf("weather key is nil")
	}
	return weatherKey, nil
}
