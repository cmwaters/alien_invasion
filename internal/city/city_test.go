package city

import (
	"testing"
)

func TestChooseRandomNeighborCity(t *testing.T) {
	thirdCity := City{
		Name: "thirdCity",
	}
	secondCity := City{
		Name: "secondCity",
	}
	firstCity := City{
		Name:  "firstCity",
		South: &thirdCity,
		West:  &secondCity,
	}
	neighborCity := firstCity.chooseRandomNeighborCity()
	if neighborCity.Name != "thirdCity" && neighborCity.Name != "secondCity" {
		t.Errorf("Failed to pick a random neigbouring city. Result: %s of type %T", neighborCity, neighborCity)
	}
}

func TestAddAndRemoveAliensFromCity(t *testing.T) {
	firstCity := City{
		Name: "firstCity",
	}
	firstCity.Aliens = append(firstCity.Aliens, 13)
	firstCity.Aliens = append(firstCity.Aliens, 25)
	if len(firstCity.Aliens) != 2 {
		t.Errorf("Failed to Add Aliens. Aliens in the city: %s", firstCity.aliensToString())
	}
	firstCity.removeAlien(13)
	if firstCity.Aliens[0] != 25 || len(firstCity.Aliens) != 1 {
		t.Errorf("Failed to remove alien 13 from city. Aliens in city: %s", firstCity.aliensToString())
	}
	firstCity.removeAlien(25)
	if len(firstCity.Aliens) != 0 {
		t.Errorf("Failed to remove alien 25 from city. Aliens in city: %s", firstCity.aliensToString())
	}
}

func TestAddAlienToACityMap(t *testing.T) {
	var cities = map[string]City{"secondCity": City{Name: "secondCity"}}
	secondCity := cities["secondCity"]
	secondCity.Aliens = append(cities["secondCity"].Aliens, 31)
	cities["secondCity"] = secondCity
	if len(cities["secondCity"].Aliens) != 1 {
		t.Errorf("Failed to add Alien. Aliens in the city: %s", cities["secondCity"].Aliens)
	}

}
