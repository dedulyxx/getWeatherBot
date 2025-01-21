package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
)

func getWeather(cityName string) string {
	appid := os.Getenv("APIKEY")
	if appid == "" {
		fmt.Printf("API KEY не установлен")
		log.Fatalf("API KEY не установлен")
	}
	// Выполнение запроса на API OpenWeatherMap
	response, err := resty.New().R().
		SetQueryParams(map[string]string{
			"q":     cityName,
			"appid": appid,
			"units": "metric",
		}).
		Get("https://api.openweathermap.org/data/2.5/weather")

	if err != nil {
		log.Fatalf("Failed to fetch weather data: %v", err)
	}

	// Проверка статуса ответа
	if response.StatusCode() == 404 {
		return "Что-то пошло не так, возможно такого города не существует, попробуй еще раз"
	}
	if response.StatusCode() != 200 {
		log.Fatalf("Error: Received status code %d", response.StatusCode())
	}

	// Структура для response
	type weatherData struct {
		Weather []struct {
			Main        string `json:"main"`
			Description string `json:"description"`
		} `json:"weather"`
		Main struct {
			Temperature float64 `json:"temp"`
			Humidity    int     `json:"humidity"`
		} `json:"main"`
	}

	var w1 weatherData

	err = json.Unmarshal(response.Body(), &w1)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("Описание погоды: %s - %s\nТемпература: %.2f°C\nВлажность: %d%%\n", w1.Weather[0].Main, w1.Weather[0].Description, w1.Main.Temperature, w1.Main.Humidity)

}
