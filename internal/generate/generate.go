// Generates a random map of cities and their neighboring cities, outputted in a .txt auxiliary_functions

package generate

import (
	"bytes"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	c "github.com/cmwaters/alien_invasion/internal/city"
	"github.com/cmwaters/alien_invasion/internal/log"
	"math/rand"
	"os"
)

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
				cities[x+sizeX*y] = &c.City{Name: cityNames[rand.Intn(len(cityNames))]}
			}
			if y > 0 { // if not the first row then link this city with the city directly above it
				cities[x+sizeX*y].North = cities[x+sizeX*(y-1)]
			}
			if x > 0 { // if not the first column then link this city with the city to the left of it which should already be created
				cities[x+sizeX*y].West = cities[(x-1)+sizeX*y]
			}
			if x < sizeX { // if not the last city  in the row then create another city to the left and link this city to it
				cities[(x+1)+sizeX*y] = &c.City{Name: cityNames[rand.Intn(len(cityNames))]}
				cities[x+sizeX*y].East = cities[(x+1)+sizeX*y]
			}
			if y < sizeY { // if not the last row in the grid then create another city below and link this city to it
				cities[x+sizeX*(y+1)] = &c.City{Name: cityNames[rand.Intn(len(cityNames))]}
				cities[x+sizeX*y].South = cities[x+sizeX*(y+1)]
			}
		}
	}
	return cities
}

func MakeOutputFile(world map[int]*c.City, fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Write("error", "Failed to read the file")
	} else {
		for _, city := range world {
			message := bytes.Buffer{}
			message.WriteString(city.Name + " - ")
			if city.North != nil {
				message.WriteString("North: " + city.North.Name)
			}
			if city.East != nil {
				message.WriteString("East: " + city.East.Name)
			}
			if city.West != nil {
				message.WriteString("West: " + city.West.Name)
			}
			if city.South != nil {
				message.WriteString("North: " + city.South.Name)
			}
			message.WriteString("\n")
			_, err := file.WriteString(message.String())
			Check(err)
		}
	}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func PrintRandomCity() {
	fmt.Println(randomdata.City())
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
