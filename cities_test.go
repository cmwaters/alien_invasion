package main

import (
	"testing"
)

func TestChooseRandomNeighborCity(t *testing.T) {
	dunedin := City{
		Name: "Dunedin",
	}
	queenstown := City{
		Name: "Queenstown",
	}
	christchurch := City{
		Name:  "Christchurch",
		South: &dunedin,
		West:  &queenstown,
	}
	neighborCity := christchurch.chooseRandomNeighborCity()
	if neighborCity.Name != "Dunedin" && neighborCity.Name != "Queenstown" {
		t.Errorf("Failed to pick a random neigbouring city. Result: %s of type %T", neighborCity, neighborCity)
	}
}

func TestRemoveAlienFromCity(t *testing.T) {
	christchurh := City{
		Name: "christchurch",
	}
	christchurh.Aliens = append(christchurh.Aliens, 13)
	christchurh.Aliens = append(christchurh.Aliens, 25)
	christchurh.removeAlien(13)
	if christchurh.Aliens[0] != 25 && len(christchurh.Aliens) != 1 {
		t.Errorf("Failed to remove alien 13 from city. Aliens in city: %s", christchurh.Aliens)
	}
}
