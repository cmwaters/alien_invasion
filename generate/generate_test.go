package generate

import (
	c "github.com/cmwaters/alien_invasion/city"
	"testing"
)

func TestMakeCityGrid(t *testing.T) {
	var cities map[string]*c.City
	cities = make(map[string]*c.City)
	cities = MakeCityGrid(4, 4)
	// test that the correct amount of cities is correct
	if len(cities) != 16 {
		t.Errorf("Generate package did not create the correct amount of cities")
	}
	// test that each of the pairings between neighbor cities is correct
	for _, city := range cities {
		if city.North != nil {
			// the city to the north of it should have this city as it's southern neighbor
			if city.North.South != city {
				t.Errorf("Incorrect north/south pairing of neighbors: %s - %s - %s", city.Name, city.North.Name, city.North.South.Name)
			}
		}
		if city.East != nil {
			// the city to the east of it should have this city as it's western neighbor
			if city.East.West != city {
				t.Errorf("Incorrect east/west pairing of neighbors: %s - %s - %s", city.Name, city.East.Name, city.East.West.Name)
			}
		}
		if city.West != nil {
			// the city to the west of it should have this city as it's eastern neighbor
			if city.West.East != city {
				t.Errorf("Incorrect west/east pairing of neighbors: %s - %s - %s", city.Name, city.West.Name, city.West.East.Name)
			}
		}
		if city.South != nil {
			// the city to the south of it should have this city as it's northern neighbor
			if city.South.North != city {
				t.Errorf("Incorrect south/north pairing of neighbors: %s - %s - %s", city.Name, city.South.Name, city.South.North.Name)
			}
		}
	}
}
