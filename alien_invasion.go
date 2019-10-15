package main

import (
	"bufio"
	"bytes"
	"fmt"
	a "github.com/cmwaters/alien_invasion/alien"
	c "github.com/cmwaters/alien_invasion/city"
	g "github.com/cmwaters/alien_invasion/general"
	"github.com/cmwaters/alien_invasion/generate"
	"github.com/cmwaters/alien_invasion/log"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var cities map[string]*c.City
var aliens map[int]*a.Alien
var WorldSize = 3

func main() {
	cities = make(map[string]*c.City)
	aliens = make(map[int]*a.Alien)
	log.Initialize("debug")
	log.Initialize("error")
	fmt.Printf("Welcome to Alien Invader\n")
	if processFile(GetFileName()) {
		injectAliens()
	} else {
		log.Write("error", "Failed to process input file")
	}
	simIterations := 0
	aliensDestroyed := false
	log.Write("debug", "Initial Simulation Configuration: "+"\n"+simulationStatus())
	for simIterations < 10000 && !aliensDestroyed {
		simIterations++
		// step through the simulation by moving the aliens and evaluating the conflicts
		step()
		// check if all aliens have been destroyed
		if len(aliens) == 0 {
			aliensDestroyed = true
		}
		if simIterations%100 == 0 { //log summary of status to the debug log
			log.Write("debug", "step: "+strconv.Itoa(simIterations)+"\n")
			log.Write("debug", simulationStatus())
		}
	}
	// convert to the correct form for the MakeOutputFile function
	generate.MakeOutputFile(cities, "output/output.txt")
}

// Steps through one iteration of the simulation. This consists of moving each Aliens to a randomly selected
// neighbor city to the one it currently resides in and then at the close checking if any of the cities contain
// two or more aliens in which both the city and the aliens inhibiting it are destroyed and removed from the map
func step() {
	// Alien moves in one of the four directions towards a new city
	for _, alien := range aliens {
		moveAlien(*alien)
	}
	// Evaluate any conflicts that may occur in a city
	for _, city := range cities {
		if len(city.Aliens) >= 2 {
			log.Write("debug", city.Name+" has been destroyed")
			fmt.Printf("%s has been destroyed by ", city.Name)
			firstTime := true
			for index, alien := range city.Aliens {
				if !firstTime {
					fmt.Printf(" and ")
				}
				firstTime = false
				fmt.Print(alien.Name())
				log.Write("debug", alien.Name()+" has been destroyed")
				delete(aliens, index)
			}
			fmt.Printf("\n")
			// delete the city and remove all relation of other cities to this city
			if cities[city.Name].North != nil {
				cities[city.Name].North.South = nil
			}
			if cities[city.Name].South != nil {
				cities[city.Name].South.North = nil
			}
			if cities[city.Name].East != nil {
				cities[city.Name].East.West = nil
			}
			if cities[city.Name].West != nil {
				cities[city.Name].West.East = nil
			}
			delete(cities, city.Name)
		}
	}
}

// Takes a txt file (as a string) and iterates over each line in search for cities to add to the cities map
// Arguments: fileName string -> the path to the file to be read
// Returns: bool -> whether it was able to succesfully read the file
func processFile(fileName string) bool {
	// open the file and check there are no errors
	file, err := os.Open(fileName)
	if err != nil {
		log.Write("error", "Failed to read the file")
		return false
	}
	// create an instance of the scanner to iterate through each line of the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x := scanner.Text()
		x = x + " "
		// for each line add the new city (or potentially cities) to the cities dictionary
		addCity(x)
	}
	// look to catch any errors in the scanning of the file
	if err := scanner.Err(); err != nil {
		log.Write("error", "Failed to scan the file")
	}
	// once complete, close the file
	err = file.Close()
	g.Check(err)
	return true
}

// The addCity function takes
func addCity(x string) {
	directions := [4]string{"north", "east", "south", "west"}
	index := 1
	breakpoint := 0
	cityFound := ""
	for index < len(x) {
		// check to see that there is a space followed by a direction which is used as the trigger to separate each
		// piece of information in a single line
		if g.NextWordIsADirection(x, index) || index == len(x)-1 { //x[index] == ' ' {
			// check to see if it is the first word in the line which will be treated as the name of the city
			if cityFound == "" {
				cityFound = x[0:index]
				// if the new city is not currently in the cities dictionary then create a new one and add it
				if _, ok := cities[x[breakpoint:index]]; !ok {
					cities[cityFound] = &c.City{
						Name: x[0:index],
					}
					city := *cities[cityFound]
					city.Aliens = make(map[int]*a.Alien)
					city.Aliens[0] = &a.Alien{
						Id:   0,
						City: city.Name,
					}
				} else {
					continue
				}
			} else {
				for _, direction := range directions {
					// check what direction followed after the previous breakpoint and therefore corresponds to this city
					if x[breakpoint+1:breakpoint+len(direction)+2] == direction+"=" {
						name := x[breakpoint+len(direction)+2 : index]
						city := cities[cityFound]
						_, ok := cities[name]
						// if this new city doesn't exist yet then create that city and add it to the dictionary
						if !ok {
							cities[name] = &c.City{Name: name}
							// also add the original city as its counterpart
							if direction == "north" {
								cities[name].South = cities[cityFound]
							} else if direction == "east" {
								cities[name].West = cities[cityFound]
							} else if direction == "west" {
								cities[name].East = cities[cityFound]
							} else if direction == "south" {
								cities[name].North = cities[cityFound]
							}
						}
						// create a pointer pointing to this city based on its appropriate direction
						if direction == "north" {
							city.North = cities[name]
						} else if direction == "east" {
							city.East = cities[name]
						} else if direction == "west" {
							city.West = cities[name]
						} else if direction == "south" {
							city.South = cities[name]
						}
						break
					}
				}
			}
			// update the new breakpoint to the current index
			breakpoint = index
		}
		index++
	}
}

func injectAliens() {
	reader := bufio.NewReader(os.Stdin)
	// create a city index so as to easily assign a city to an alien based from a randomly generated int
	var cityIndex map[int]string
	cityIndex = make(map[int]string)
	index := 0
	for _, city := range cities {
		cityIndex[index] = city.Name
		index++
	}
	distributedAliens := false
	for !distributedAliens {
		// Ask user for the amount of Aliens to enter the simulation
		fmt.Print("Enter the amount of Aliens to enter the simulation: ")
		amountOfAliensString, _ := reader.ReadString('\n')
		amountOfAliensString = strings.TrimSuffix(amountOfAliensString, "\n")
		amountOfAliens, err := strconv.Atoi(amountOfAliensString)
		if err != nil {
			log.Write("error", "Value provided was unable to be converted to a string.")
		} else {
			// Create each alien
			for index := 0; index < amountOfAliens; index++ {
				// Find a random index to pick out a random city to drop the alien in
				randNum := rand.Intn(len(cities))
				// Create the alien and add it to the aliens map as a pointer
				aliens[index] = &a.Alien{
					Id:   index,
					City: cityIndex[randNum],
				}
				// Reference the same pointer to the alien within the appropriate city in the cities map
				alien := aliens[index]
				if cities[cityIndex[randNum]].Aliens == nil {
					cities[cityIndex[randNum]].Aliens = make(map[int]*a.Alien)
					log.Write("error", "Reinitialized the Alien Map whilst injecting")
				}
				cities[cityIndex[randNum]].Aliens[index] = alien
			}
			distributedAliens = true
		}
	}
}

// moves an alien (denoted as an int) from its current city to a neighboring city
func moveAlien(alien a.Alien) {
	if alien.City == "" {
		log.Write("error", "Alien "+strconv.Itoa(alien.Id)+" has no city assigned to it")
	} else if cities[alien.City] == nil {
		log.Write("error", "The city, "+alien.City+" does not exist")
	} else if !cities[alien.City].AlienPresentInCity(alien) {
		log.Write("error", "Alien "+strconv.Itoa(alien.Id)+" is assumed to be present in a city that "+
			"it isn't in")
	} else {
		oldCity := cities[alien.City]
		newCity := oldCity.ChooseRandomNeighborCity()
		// check that the alien has a new city to move to
		if !(newCity.Name == oldCity.Name) {

			log.Write("debug", alien.Name()+" has moved from Old City: "+oldCity.Name+" to new City: "+newCity.Name)
			delete(oldCity.Aliens, alien.Id)
			if newCity.Aliens == nil {
				newCity.Aliens = make(map[int]*a.Alien)
				log.Write("error", "Reinitializing the Alien Map")
			}
			newCity.Aliens[alien.Id] = &alien
			aliens[alien.Id].City = newCity.Name
		}
	}
}

func simulationStatus() string {
	message := bytes.Buffer{}
	for _, city := range cities {
		message.WriteString(city.Name + ":")
		for alien := range city.Aliens {
			message.WriteString(" Alien " + strconv.Itoa(alien))
		}
		message.WriteString("\n")
	}
	for _, alien := range aliens {
		message.WriteString(alien.Name() + ": " + alien.City + "\n")
	}
	return message.String()
}

func GetFileName() string {
	reader := bufio.NewReader(os.Stdin)
	files, err := ioutil.ReadDir("maps/")
	g.Check(err)
	if len(files) >= 1 {
		for _, file := range files {
			fmt.Println(file.Name())
		}
		for i := 0; i < 10; i++ {
			fmt.Print("Please enter the name of the map you wish to simulate (i.e test_map) or press G to generate one: ")
			mapName, _ := reader.ReadString('\n')
			mapName = strings.TrimSuffix(mapName, "\n")
			if mapName == "G" || mapName == "g" {
				generateMap()
				return GetFileName()
			}
			for _, file := range files {
				if mapName+".txt" == file.Name() {
					fmt.Printf("Running simulation with map: %s \n", mapName+".txt")
					return "maps/" + mapName + ".txt"
				}
			}
		}
		fmt.Print("Closing Application")
		return ""
	} else {
		// If there are no maps in the directory then request to generate one
		fmt.Print("You have no maps in the maps directory. Would you like to generate one (Y/N): ")
		output, _ := reader.ReadString('\n')
		output = strings.TrimSuffix(output, "\n")
		if output == "Y" || output == "y" || output == "yes" || output == "Yes" {
			generateMap()
			return GetFileName()
		} else {
			return ""
		}
	}
}

func generateMap() {
	fileNameFound := false
	var fileName string
	i := 1
	for !fileNameFound {
		fileName = "maps/generated_map_" + strconv.Itoa(i) + ".txt"
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			fileNameFound = true
			break
		}
		i++
	}
	generate.MakeOutputFile(generate.MakeCityGrid(WorldSize, WorldSize), fileName)
}
