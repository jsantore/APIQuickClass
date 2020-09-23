package main

type WeatherData struct{
	Title			string `json:title`
	LocationType 	string `json:location_type`
	ID 				string `json:woeid`
	Location 		string `json:latt_long`
}
