package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	config "./config"
)

// Gets the Zip Code from user input and returns the zip code
func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin) // Create a scanner object using the Input from the OS
	fmt.Println("Please enter the zip code:")
	scanner.Scan()        // Get user input
	zip := scanner.Text() // Get the text from the input
	// Short validation for the zip to make sure it is 5 characters
	// FEATURE ADDITION? Make sure it is an int too
	if len(zip) != 5 {
		fmt.Println("Please enter a valid zip code")
		zip = getUserInput()
	}
	return zip
}

// Calls to the API and stores response information in a struct object return the struct variable
func getAPI(apiString string) Weather {
	// Create a var of the struct to hold information
	var newWeather Weather
	// Use http.Get to call the API
	response, err := http.Get(apiString)
	// Make sure there is no error
	if err != nil {
		fmt.Println(err)
	} else {
		// Decode the JSON and save it in the variable we created
		jsonErr := json.NewDecoder(response.Body).Decode(&newWeather)
		if jsonErr != nil {
			fmt.Println(jsonErr)
		}
	}
	return newWeather
}

// Basic output for the information just to make sure everything is working
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
	// Set location
	location := getUserInput()
	// Format the string for the API call with my private key and the zip code provided by the user
	apiString := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?zip=%s&appid=%s&units=imperial", location, config.GetAPIKey())
	// Get the weather from the API
	weatherForLoc := getAPI(apiString)
	// Print the weather
	outputWeather(weatherForLoc)
}
