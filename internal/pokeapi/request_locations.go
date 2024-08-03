package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

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
