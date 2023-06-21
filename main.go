package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
)

// this is giving me a headache man idek if im doing this right
type Weather struct{
	Location struct{
		Name string `json:"name"`
		Localtime string `json: localtime`
	} `json:"location"`
	Current struct{
		TempC float64 `json:"temp_c"`
		Condition struct {
			Text string `json: "text"`
		} `json: condition`
	} `json:"current"`
}

func main(){
	fmt.Println("Hey there! Want a raincheck on a city?")
	for {
	var userInput string
	fmt.Print("Drop the place: ")
	fmt.Scan(&userInput)

	res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=8bd70fc3fa2e4593815201535231606&q="+userInput+"&aqi=no")
	if err != nil{
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200{
		panic("Sorry! Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil{
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current := weather.Location, weather.Current

	var output = fmt.Sprintf(
		"%s \t\t %s \t %.0fC (%s)\n",
		location.Localtime[len(location.Localtime)-5:len(location.Localtime)],
		location.Name,
		current.TempC,
		current.Condition.Text,
	)

	if current.TempC > 20 {
		color.Red(output)
	} else if current.TempC > 15 {
		color.Yellow(output)
	} else if current.TempC > 10 {
		color.Green(output)
	} else if current.TempC > 5 {
		color.Blue(output)
	} else if current.TempC > 0 {
		color.Cyan(output)
	} else {
		color.White(output)
	}

	var userInputLoop string
	fmt.Print("Want to check another place? [y/n]: ")
	fmt.Scan(&userInputLoop)
	if userInputLoop != "y"{
		break
	}
}
}