package main

import (
	"net/http"
)

func main(){
	res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=8bd70fc3fa2e4593815201535231606&q=Toronto&aqi=no")
	if err != nil{
		panic(err)
	}
}