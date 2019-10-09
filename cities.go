// Used to create cities in which aliens can move around and potentially destroy
package main

import "fmt"

type City struct {
	Name  string // city name must be unique
	North *City
	East  *City
	West  *City
	South *City
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
	fmt.Printf("\n")
}
