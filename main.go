package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=8bd70fc3fa2e4593815201535231606&q=Toronto&aqi=no")
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
	fmt.Println(weather)
}