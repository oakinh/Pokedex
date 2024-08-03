package main

import (
	"errors"
	"fmt"

	"github.com/oakinh/pokedex/internal/pokeapi"
)

var mapPosition int = 1

func commandMap() error {
	for i := mapPosition; i < mapPosition+20; i++ {
		locationName := pokeapi.RequestLocations(i)
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
		locationName := pokeapi.RequestLocations(i)
		fmt.Println(locationName)

	}
	mapPosition += 20
	return nil
}
