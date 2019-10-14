// Generates a random map of cities and their neighboring cities, outputted in a .txt general
package generate

import (
	"fmt"
	c "github.com/cmwaters/alien_invasion/internal/city"
	. "github.com/cmwaters/alien_invasion/internal/general"
	"github.com/cmwaters/alien_invasion/internal/log"
	"math/rand"
	"os"
)

var usedCityNames []string

// Creates a grid of cities that is sizeX wide by sizeY high
func MakeCityGrid(sizeX int, sizeY int) map[int]*c.City {
	if sizeX*sizeY > 100 { // Can't have more than 100 unique cities
		sizeX = 10
		sizeY = 10
	}
	var cities map[int]*c.City
	cities = make(map[int]*c.City)
	// loop along both the x and y axes
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			if cities[x+sizeX*y] == nil { // should only run for creating the first city (at the top left corner of the grid)
				cities[x+sizeX*y] = &c.City{Name: randomUniqueCity()}
			}
			if y > 0 { // if not the first row then link this city with the city directly above it
				cities[x+sizeX*y].North = cities[x+sizeX*(y-1)]
			}
			if x > 0 { // if not the first column then link this city with the city to the left of it which should already be created
				cities[x+sizeX*y].West = cities[(x-1)+sizeX*y]
			}
			if x < sizeX-1 { // if not the last city  in the row then create another city to the left and link this city to it
				if y == 0 {
					cities[(x+1)+sizeX*y] = &c.City{Name: randomUniqueCity()}
				}
				cities[x+sizeX*y].East = cities[(x+1)+sizeX*y]
			}
			if y < sizeY-1 { // if not the last row in the grid then create another city below and link this city to it
				cities[x+sizeX*(y+1)] = &c.City{Name: randomUniqueCity()}
				cities[x+sizeX*y].South = cities[x+sizeX*(y+1)]
			}
		}
	}
	return cities
}

func MakeOutputFile(world map[int]*c.City, fileName string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Write("error", "Failed to write to the file: "+fileName)
	} else {
		for _, city := range world {
			_, err := file.WriteString(city.String())
			Check(err)
		}
	}
	err = file.Close()
	Check(err)
	fmt.Println("Successfully created a new map under the path: " + fileName)
}

func randomUniqueCity() string {
	uniqueCityFound := false
	var cityName string
	for !uniqueCityFound {
		cityName = cityNames[rand.Intn(len(cityNames))]
		if !Contains(usedCityNames, cityName) {
			uniqueCityFound = true
		}
	}
	usedCityNames = append(usedCityNames, cityName)
	return cityName
}

// source from https://simplemaps.com/data/de-cities
var cityNames = []string{
	"Berlin",
	"Stuttgart",
	"Frankfurt",
	"Mannheim",
	"Hamburg",
	"Essen",
	"Duisburg",
	"Munich",
	"Düsseldorf",
	"Cologne",
	"Wuppertal",
	"Saarbrücken",
	"Marienberg",
	"Bremen",
	"Hannover",
	"Bonn",
	"Dresden",
	"Wiesbaden",
	"Dortmund",
	"Leipzig",
	"Heidelberg",
	"Karlsruhe",
	"Augsburg",
	"Bielefeld",
	"Koblenz",
	"Altchemnitz",
	"Kassel",
	"Münster",
	"Kiel",
	"Freiburg",
	"Braunschweig",
	"Fürth",
	"Lübeck",
	"Osnabrück",
	"Magdeburg",
	"Potsdam",
	"Erfurt",
	"Rostock",
	"Mainz",
	"Ulm",
	"Würzburg",
	"Oldenburg",
	"Regensburg",
	"Ingolstadt",
	"Göttingen",
	"Bremerhaven",
	"Cottbus",
	"Jena",
	"Gera",
	"Flensburg",
	"Schwerin",
	"Rosenheim",
	"Gießen",
	"Stralsund",
	"Coburg",
	"Hofeck",
	"Emden",
	"Detmold",
	"Meißen",
	"Kitzingen",
	"Dingolfing",
	"Heppenheim",
	"Torgau",
	"Hanau",
	"Husum",
	"Schwandorf",
	"Bitburg",
	"Cham",
	"Traunstein",
	"Lüchow",
	"Gifhorn",
	"Biberach",
	"Bad Reichenhall",
	"Künzelsau",
	"Weißenburg",
	"Regen",
	"Nuremberg",
	"Aurich",
	"Nordhorn",
	"Aichach",
	"Marburg",
	"Görlitz",
	"Vechta",
	"Trier",
	"Pirmasens",
	"Pirna",
	"Neustadt",
	"Beeskow",
	"Westerstede",
	"Verden",
	"Worms",
	"Düren",
	"Landsberg",
	"Ludwigsburg",
	"Meiningen",
	"Siegen",
	"Deggendorf",
	"Peine",
	"Frankfurt (Oder)",
	"Nienburg",
	"Brake",
	"Memmingen",
	"Kirchheimbolanden",
	"Tauberbischofsheim",
	"Emmendingen",
	"Warendorf",
	"Bad Segeberg",
	"Rotenburg",
	"Kronach",
	"Darmstadt",
	"Mindelheim",
	"Bergheim",
	"Donauwörth",
	"Korbach",
}
