package main

import (
	"bufio"
	"fmt"
	"os"

	config "./config"
)

func getUserInput() string {
	location := ""
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the city you live in:")
	scanner.Scan()
	city := scanner.Text()
	if city == "" {
		fmt.Println("Please enter a valid name")
		location = getUserInput()
		return location
	}
	fmt.Println("Please enter the State Code (EX: MA, NY, CA):")
	scanner.Scan()
	state := scanner.Text()
	if state == "" || len(state) != 2 {
		fmt.Println("Please enter a valid name")
		location = getUserInput()
		return location
	}
	fmt.Println("Please enter the 2 letter Country Code (EX: US, UK):")
	scanner.Scan()
	country := scanner.Text()
	if country == "" || len(country) != 2 {
		fmt.Println("Please enter a valid name")
		location = getUserInput()
		return location
	}
	location = city + "," + state + "," + country
	return location
}

func main() {
	fmt.Println("The api key is", config.GetAPIKey())
	location := getUserInput()
	apiString := fmt.Sprintf("api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=imperial", location, config.GetAPIKey())
	fmt.Println("The city entered was", location, "and the api string is", apiString)
}
