package main

import (
	"github.com/cmwaters/alien_invasion/internal/city"
	"testing"
)

func TestAlienMovesFromOneCityToAnother(t *testing.T) {
	firstCity := city.City{Name: "firstCity"}
	secondCity := city.City{Name: "secondCity"}
	thirdCity := city.City{Name: "thirdCity"}
	firstCity.Aliens = append(firstCity.Aliens, 7)
	// function should not move an alien if an alien has no city to go to
	oldCity, newCity := moveAlien(7, firstCity)
	if newCity.Name != oldCity.Name || len(newCity.Aliens) != 1 {
		t.Errorf("Alien has failed to stay trapped in its city. Result %s : %s", newCity.Name, newCity.aliensToString())
	}
	// function should move alien from first city to second city
	firstCity.West = &secondCity
	secondCity.East = &firstCity
	oldCity, newCity = moveAlien(7, firstCity)
	if len(oldCity.Aliens) != 0 || newCity.Aliens[0] != 7 {
		t.Errorf("Failed to Move alien. There should be 1 alien in secondCity. There is %d. There should be 0 "+
			"aliens in firstCity. There is %d.", len(newCity.Aliens), len(oldCity.Aliens))
	}
	// function should return the original city if the alien is not in the city (an error is also produced)
	oldCity, newCity = moveAlien(13, secondCity)
	if oldCity.Name != secondCity.Name || newCity.Name != secondCity.Name {
		t.Errorf("Alien which doesn't exist has moved")
	}
	// function to check that multiple aliens can move across to one city
	thirdCity.North = &firstCity
	firstCity.South = &thirdCity
	thirdCity.Aliens = append(thirdCity.Aliens, 11)
	thirdCity.Aliens = append(thirdCity.Aliens, 4)
	oldCity, newCity = moveAlien(11, thirdCity)
	thirdCity = oldCity
	firstCity = newCity
	oldCity, newCity = moveAlien(4, thirdCity)
	thirdCity = oldCity
	firstCity = newCity
	if len(thirdCity.Aliens) != 0 || len(firstCity.Aliens) != 3 {
		t.Errorf("Aliens haven't moved succesfully across from third city to first city. Respective Cities: %s;"+
			" %s", thirdCity.aliensToString(), firstCity.aliensToString())
	}

}

func TestAliensDestroyEachOtherAndCity(t *testing.T) {

}
