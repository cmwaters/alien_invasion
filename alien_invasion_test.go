package main

import (
	a "github.com/cmwaters/alien_invasion/alien"
	c "github.com/cmwaters/alien_invasion/city"
	"testing"
)

func TestAlienMovesFromOneCityToAnother(t *testing.T) {
	cities = make(map[string]*c.City)
	aliens = make(map[int]*a.Alien)
	cities["firstCity"] = &c.City{Name: "firstCity"}
	cities["secondCity"] = &c.City{Name: "secondCity"}
	cities["thirdCity"] = &c.City{Name: "thirdCity"}
	aliens[7] = &a.Alien{
		Id:   7,
		City: cities["firstCity"].Name,
	}
	cities["firstCity"].Aliens = make(map[int]*a.Alien)
	cities["firstCity"].Aliens[7] = aliens[7]
	// function should not move an alien if an alien has no city to go to
	moveAlien(*aliens[7])
	if aliens[7].City != cities["firstCity"].Name || len(cities["firstCity"].Aliens) != 1 {
		t.Errorf("Alien has failed to stay trapped in its city. Result: \n" + simulationStatus())
	}
	// function should move alien from first city to second city
	cities["firstCity"].West = cities["secondCity"]
	cities["secondCity"].East = cities["firstCity"]
	moveAlien(*aliens[7])
	if len(cities["firstCity"].Aliens) != 0 || len(cities["secondCity"].Aliens) != 1 {
		t.Errorf("Failed to Move alien. There should be 1 alien in secondCity. There is %d. There should be 0 "+
			"aliens in firstCity. There is %d.", len(cities["secondCity"].Aliens), len(cities["firstCity"].Aliens))
	}
	// function to check that multiple aliens can move across to one city
	cities["thirdCity"].North = cities["secondCity"]
	cities["secondCity"].South = cities["thirdCity"]
	aliens[11] = &a.Alien{
		Id:   11,
		City: "thirdCity",
	}
	aliens[4] = &a.Alien{
		Id:   4,
		City: "thirdCity",
	}
	cities["thirdCity"].Aliens = make(map[int]*a.Alien)
	cities["thirdCity"].Aliens[11] = aliens[11]
	cities["thirdCity"].Aliens[4] = aliens[4]
	moveAlien(*aliens[11])
	moveAlien(*aliens[4])
	if len(cities["thirdCity"].Aliens) != 0 || len(cities["secondCity"].Aliens) != 3 {
		t.Errorf("Aliens haven't moved succesfully across from third city to second city. Status: \n" +
			simulationStatus())
	}

}

func TestAliensDestroyEachOtherAndCity(t *testing.T) {

}
