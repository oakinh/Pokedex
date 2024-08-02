package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

var mapPosition int = 1

func requestLocations(id int) string {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d/", id)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if response.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", response.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	//dat := []byte(body)
	locationArea := LocationArea{}
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		fmt.Println(err)
	}
	if locationArea.Name != nil && *locationArea.Name != "" {
		return *locationArea.Name
	}

	return "Unknown Location"
}

func commandMap() error {
	for i := mapPosition; i < mapPosition+20; i++ {
		locationName := requestLocations(i)
		fmt.Println(locationName)

	}
	mapPosition += 20
	return nil
}

func commandMapb() error {
	if mapPosition-40 < 1 {
		return errors.New("you cannot go back further than the first result. Please continue in the map before attempting to go back")
	}
	mapPosition -= 40
	for i := mapPosition; i < mapPosition+20; i++ {
		locationName := requestLocations(i)
		fmt.Println(locationName)

	}
	mapPosition += 20
	return nil
}
