// Used to create cities in which aliens can move around and potentially destroy
package city

import (
	"bytes"
	"fmt"
	a "github.com/cmwaters/alien_invasion/internal/alien"
	"math/rand"
	"strconv"
)

type City struct {
	Name   string // city name must be unique
	North  *City
	East   *City
	West   *City
	South  *City
	Aliens map[int]*a.Alien
}

func (city City) ChooseRandomNeighborCity() *City {
	var cityRoulette map[int]*City
	cityRoulette = make(map[int]*City)
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
	neighborCity := cityRoulette[rand.Intn(len(cityRoulette))]
	return neighborCity

}

func (city City) AlienPresentInCity(alien a.Alien) bool {
	for alienName, _ := range city.Aliens {
		if alienName == alien.Id {
			return true
		}
	}
	return false
}

func (city City) Print() {
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
	for _, alien := range city.Aliens {
		fmt.Printf("Alien %d ", alien)
	}
	fmt.Printf("\n")
}

func (city City) AliensToString() string {
	message := bytes.Buffer{}
	for id, _ := range city.Aliens {
		message.WriteString("Alien " + strconv.Itoa(id) + ", ")
	}
	return message.String()
}
