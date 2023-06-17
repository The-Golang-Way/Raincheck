package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
		Feelslike_C float64 `json: "feelslike_c"`
	} `json:"current"`
}

func main(){
	userInput := "Toronto"
	if len(os.Args) >= 2{
		userInput = os.Args[1]
	}

	res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=8bd70fc3fa2e4593815201535231606&q="+ userInput+"&aqi=no")
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

	fmt.Printf(
		"%s \t\t %s \t %.0fC (%s)\n",
		location.Localtime[len(location.Localtime)-5:len(location.Localtime)],
		location.Name,
		current.TempC,
		current.Condition.Text,
	)
}