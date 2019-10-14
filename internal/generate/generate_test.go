package generate

import (
	"bytes"
	c "github.com/cmwaters/alien_invasion/internal/city"
	"testing"
)

func TestMakeCityGrid(t *testing.T) {
	var cities map[int]*c.City
	cities = make(map[int]*c.City)
	cities = MakeCityGrid(2, 2)
	if len(cities) != 10 {
		message := bytes.Buffer{}
		for _, city := range cities {
			message.WriteString(city.Name + " -")
			if city.North != nil {
				message.WriteString(" North: " + city.North.Name)
			}
			if city.East != nil {
				message.WriteString(" East: " + city.East.Name)
			}
			if city.West != nil {
				message.WriteString(" West: " + city.West.Name)
			}
			if city.South != nil {
				message.WriteString(" North: " + city.South.Name)
			}
			message.WriteString("\n")
		}
		t.Errorf(message.String())
	}
}
