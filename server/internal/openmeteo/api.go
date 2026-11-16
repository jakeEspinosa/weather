package openmeteo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type WeatherResponse struct {
	Current CurrentData `json:"current"`
}

type CurrentData struct {
	Time          string  `json:"time"`
	Temperature2m float64 `json:"temperature_2m"`
}

func FetchWeather() {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get("https://api.open-meteo.com/v1/forecast?latitude=39.7294&longitude=-104.8319&current=temperature_2m,relative_humidity_2m,precipitation,rain,showers,snowfall,surface_pressure&timezone=America%2FDenver&forecast_days=1&wind_speed_unit=mph&temperature_unit=fahrenheit&precipitation_unit=inch")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	var weather WeatherResponse
	weatherErr := json.NewDecoder(resp.Body).Decode(&weather)
	if weatherErr != nil {
		fmt.Printf("Invalid")
	}

	fmt.Println(weather)
}
