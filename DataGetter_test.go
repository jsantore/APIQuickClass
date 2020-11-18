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
	data, err := ioutil.ReadAll(answer.Body)
	if err != nil {
		t.Error("error reading")
	}
	if len(data) < 1 {
		t.Error("there was no data")
	}
}
