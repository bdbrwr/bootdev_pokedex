package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MapArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type MapAreaResults struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []MapArea `json:"results"`
}

var nextMap string = "https://pokeapi.co/api/v2/location-area/"
var prevMap string = ""

func commandMap() error {
	res, err := http.Get(nextMap)
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	results := MapAreaResults{}
	err = json.Unmarshal(body, &results)
	if err != nil {
		fmt.Println(err)
	}

	nextMap = results.Next
	prevMap = results.Previous

	areas := results.Results

	for _, area := range areas {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb() error {
	if prevMap == "" {
		return fmt.Errorf("not yet on second page, cannot call back")
	}

	res, err := http.Get(prevMap)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	results := MapAreaResults{}
	err = json.Unmarshal(body, &results)
	if err != nil {
		fmt.Println(err)
	}

	nextMap = results.Next
	prevMap = results.Previous

	areas := results.Results

	for _, area := range areas {
		fmt.Println(area.Name)
	}
	return nil
}
