// Used to create cities in which aliens can move around and potentially destroy
package main

import (
	"fmt"
	"math/rand"
)

type City struct {
	Name   string // city name must be unique
	North  *City
	East   *City
	West   *City
	South  *City
	Aliens []int
}

func (city City) print() {
	fmt.Printf("Name: %s, ", city.Name)
	if city.North != nil {
		fmt.Printf("North: %s, ", city.North.Name)
	}
	if city.East != nil {
		fmt.Printf("East: %s, ", city.East.Name)
	}
	if city.West != nil {
		fmt.Printf("West: %s, ", city.West.Name)
	}
	if city.South != nil {
		fmt.Printf("South: %s, ", city.South.Name)
	}
	for alien := range city.Aliens {
		fmt.Printf("Alien %d", alien)
	}
	fmt.Printf("\n")
}

func (city City) chooseRandomNeighborCity() City {
	var cityRoulette map[int]City
	cityRoulette = make(map[int]City)
	counter := 0
	if city.North != nil {
		cityRoulette[counter] = *city.North
		counter++
	}
	if city.East != nil {
		cityRoulette[counter] = *city.East
		counter++
	}
	if city.West != nil {
		cityRoulette[counter] = *city.West
		counter++
	}
	if city.South != nil {
		cityRoulette[counter] = *city.South
	}
	neighborCity := cityRoulette[rand.Intn(len(cityRoulette))]
	return neighborCity

}

func (city City) removeAlien(alienId int) {
	var i int
	for index, alien := range city.Aliens {
		if alien == alienId {
			i = index
			break
		}
	}
	if len(city.Aliens) > 1 {
		city.Aliens[i] = city.Aliens[len(city.Aliens)-1] // overwrite the last element to index i.
		//city.Aliens[len(city.Aliens)-1] = ""   // Erase last element (write zero value).
		city.Aliens = city.Aliens[:len(city.Aliens)-1] // truncate the array
	} else {
		city.Aliens = nil
	}
}
