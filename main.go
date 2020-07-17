package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	config "./config"
)

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the zip code:")
	scanner.Scan()
	zip := scanner.Text()
	if len(zip) != 5 {
		fmt.Println("Please enter a valid zip code")
		zip = getUserInput()
	}
	return zip
}

func getAPI(apiString string) Weather {
	var newWeather Weather
	response, err := http.Get(apiString)
	if err != nil {
		fmt.Println(err)
	} else {
		jsonErr := json.NewDecoder(response.Body).Decode(&newWeather)
		if jsonErr != nil {
			fmt.Println(jsonErr)
		}
	}
	return newWeather
}

func outputWeather(weatherForLoc Weather) {
	fmt.Println("Here is the weather for", weatherForLoc.Name)
	fmt.Println("The temperature is", weatherForLoc.Main.Temp)
}

// Weather => this is a main struct for the weather
type Weather struct {
	Name string      `json:"name"`
	Main WeatherInfo `json:"main"`
}

// WeatherInfo => Looks at the main object with json being passed into it
type WeatherInfo struct {
	Temp float64 `json:"temp"`
}

func main() {
	location := getUserInput()
	apiString := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?zip=%s&appid=%s&units=imperial", location, config.GetAPIKey())
	weatherForLoc := getAPI(apiString)
	outputWeather(weatherForLoc)
}
