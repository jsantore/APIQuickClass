package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetWeatherData() *http.Response {
	response, err := http.Get("https://www.metaweather.com/api/location/search/?query=london")
	if err != nil {
		log.Fatal(err)
	}

	//defer response.Body.Close()
	//if response.StatusCode != 200{
	//	log.Fatal("Didn't get 200")
	//}
	return response
}

func main() {
	var data [2]WeatherData
	response := GetWeatherData()
	defer response.Body.Close()
	rawData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(rawData, &data)

	fmt.Println(data[0])

}
