// Package to create the city struct and respective functions
package city

import (
	"bytes"
	a "github.com/cmwaters/alien_invasion/alien"
	"math/rand"
)

type City struct {
	Name   string // city name must be unique
	North  *City
	East   *City
	West   *City
	South  *City
	Aliens map[int]*a.Alien
}

// Adds all possible neighboring cities to a list and selects a random one. Returning the pointer address to that city.
func (city City) ChooseRandomNeighborCity() *City {
	cityRoulette := make(map[int]*City)
	counter := 0
	if city.North != nil {
		cityRoulette[counter] = city.North
		counter++
	}
	if city.East != nil {
		cityRoulette[counter] = city.East
		counter++
	}
	if city.West != nil {
		cityRoulette[counter] = city.West
		counter++
	}
	if city.South != nil {
		cityRoulette[counter] = city.South
	}
	if len(cityRoulette) == 0 {
		return &city
	}
	return cityRoulette[rand.Intn(len(cityRoulette))]
}

// Checks if a particular alien is present in a city or not. Is used when an alien is moving between cities
func (city City) AlienPresentInCity(alien a.Alien) bool {
	for alienName := range city.Aliens {
		if alienName == alien.Id {
			return true
		}
	}
	return false
}

// condenses all city related information of that struct to a string to be outputted back to a file
func (city City) String() string {
	message := bytes.Buffer{}
	message.WriteString(city.Name)
	if city.North != nil {
		message.WriteString(" north=" + city.North.Name)
	}
	if city.East != nil {
		message.WriteString(" east=" + city.East.Name)
	}
	if city.West != nil {
		message.WriteString(" west=" + city.West.Name)
	}
	if city.South != nil {
		message.WriteString(" south=" + city.South.Name)
	}
	message.WriteString("\n")
	return message.String()
}

func (city City) AliensToString() string {
	message := bytes.Buffer{}
	for _, alien := range city.Aliens {
		message.WriteString("Aliens: " + alien.Name() + " ")
	}
	return message.String()
}
