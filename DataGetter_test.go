package main

import (
	"io/ioutil"
	"testing"
)

func TestGetWeatherData(t *testing.T) {
	answer := GetWeatherData()
	if answer.StatusCode != 200 {
		t.Error("Got bad Status")
	}
	if 1 > 2 {

	}
	data, err := ioutil.ReadAll(answer.Body)
	if err != nil {
		t.Error("error reading")
	}
	//trying stuff out
	if len(data) < 1 {
		t.Error("there was no data")
	}
}
