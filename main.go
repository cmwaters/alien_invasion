package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var cities map[string]City

func main() {
	fmt.Printf("Welcome to Alien Invader\n")
	//reader := bufio.NewReader(os.Stdin)
	//_, err := ioutil.ReadDir("maps/")
	//if err != nil {
	//	// If there are no maps in the directory then request to generate one
	//	fmt.Print("You have no maps in the maps directory. Would you like to generate one (Y/N): ")
	//	// TODO: Make sure a map is available to be read and allow for one to automatically be generated
	//} else {
	//	// Else ask the user to select a map for the simulation
	//	fmt.Print("Please enter the name of the map you wish to simulate (i.e test_map): ")
	//	// TODO: Add handling of errors in user input here
	//}
	//mapName, _ := reader.ReadString('\n')
	//mapName = strings.TrimSuffix(mapName, "\n")

	cities = make(map[string]City)

	processFile("maps/test_map.txt")
	//processFile("maps/" + mapName + ".txt")

	christchurch := City{
		Name: "Christchurch",
	}

	dunedin := City{
		Name: "Dunedin",
	}

	dunedin.North = &christchurch

	fmt.Print("Hello")
	fmt.Printf("Name: %s \n", dunedin.North.Name)

	dunedin.print()

	for _, city := range cities {
		city.print()
	}

}

func processFile(fileName string) bool {
	// open file and check there are no errors
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Print("Failed to read file")
		log.Fatal(err)
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
		log.Fatal(err)
	}
	// once complete, close the file
	file.Close()
	return true
}

func addCity(x string) {
	directions := [4]string{"north", "east", "south", "west"}
	index := 1
	breakpoint := 0
	cityFound := ""
	for index < len(x) {
		// check to see that there is a space followed by a direction which is used as the trigger to separate each
		// piece of information in a single line
		if nextWordIsADirection(x, index) || index == len(x)-1 { //x[index] == ' ' {
			// check to see if it is the first word in the line which will be treated as the name of the city
			if cityFound == "" {
				cityFound = x[0:index]
				// if the new city is not currently in the cities dictionary then create a new one and add it
				if _, ok := cities[x[breakpoint:index]]; !ok {
					city := City{
						Name: x[0:index],
					}
					cities[city.Name] = city
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
							newCity := City{Name: name}
							// also add the original city as its counterpart
							if direction == "north" {
								newCity.South = &city
							} else if direction == "east" {
								newCity.West = &city
							} else if direction == "west" {
								newCity.East = &city
							} else if direction == "south" {
								newCity.North = &city
							}
							cities[name] = newCity
						}
						elem := cities[name]
						// create a pointer pointing to this city based on its appropriate direction
						if direction == "north" {
							city.North = &elem
						} else if direction == "east" {
							city.East = &elem
						} else if direction == "west" {
							city.West = &elem
						} else if direction == "south" {
							city.South = &elem
						}
						cities[cityFound] = city
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

func nextWordIsADirection(sentence string, index int) bool {
	directions := [4]string{"north", "east", "south", "west"}
	for _, direction := range directions {
		// checks to see that the comparison will not exceed the size of the sentence string
		if index+len(direction)+1 <= len(sentence) {
			// check that the following characters are a space and the possible directions
			if sentence[index:index+len(direction)+1] == " "+direction {
				return true
			}
		}
	}
	return false
}
