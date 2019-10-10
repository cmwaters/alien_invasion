package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var cities map[string]City
var aliens map[int]City

func main() {
	cities = make(map[string]City)
	aliens = make(map[int]City)
	initializeLog()
	writeLog("SIMULATION START")
	fmt.Printf("Welcome to Alien Invader\n")
	processFile(getFileName())
	injectAliens()
	for _, city := range cities {
		city.print()
	}
	simIterations := 0
	aliensDestroyed := false
	for simIterations < 10000 && !aliensDestroyed {
		// step through the simulation by moving the aliens and evaluating the conflicts
		step()
		simIterations++
		// check if all aliens have been destroyed
		if len(aliens) == 0 {
			aliensDestroyed = true
		}
		writeLog("Iteration " + strconv.Itoa(simIterations)) // + " Summary: Cities: " + len(cities) + ", Aliens: " + len(aliens))
	}
	fmt.Print(len(aliens))
	fmt.Print(simIterations)

}

func step() {
	// Alien moves in one of the four directions towards a new city
	for alien, city := range aliens {
		priorCity := aliens[alien]
		priorCity.removeAlien(alien)
		newCity := city.chooseRandomNeighborCity()
		aliens[alien] = newCity
	}
	// Evaluate any conflicts that may occur in a city
	for _, city := range cities {
		if len(city.Aliens) >= 2 {
			fmt.Printf("%s has been destroyed by alien ", city.Name)
			for index, alienId := range city.Aliens {
				fmt.Printf("%s", alienId)
				delete(aliens, alienId)
				if index < len(city.Aliens)-1 {
					fmt.Printf(" and alien ")
				}
			}
			delete(cities, city.Name)
		}
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

func getFileName() string {
	reader := bufio.NewReader(os.Stdin)
	files, err := ioutil.ReadDir("maps/")
	if err != nil {
		// If there are no maps in the directory then request to generate one
		fmt.Print("You have no maps in the maps directory. Would you like to generate one (Y/N): ")
		// TODO: Make sure a map is available to be read and allow for one to automatically be generated
		return ""
	} else if len(files) > 1 {
		for _, file := range files {
			fmt.Println(file.Name())
		}
		for i := 0; i < 10; i++ {
			fmt.Print("Please enter the name of the map you wish to simulate (i.e test_map): ")
			mapName, _ := reader.ReadString('\n')
			mapName = strings.TrimSuffix(mapName, "\n")
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
		fmt.Printf("Running simulation with map: %s \n", files[0].Name())
		return "maps/" + files[0].Name()
	}

}

func injectAliens() {
	reader := bufio.NewReader(os.Stdin)
	var cityIndex map[int]City
	cityIndex = make(map[int]City)
	index := 0
	for _, city := range cities {
		cityIndex[index] = city
		index++
	}
	distributedAliens := false
	for !distributedAliens {
		fmt.Print("Enter the amount of Aliens to enter the simulation: ")
		amountOfAliensString, _ := reader.ReadString('\n')
		amountOfAliensString = strings.TrimSuffix(amountOfAliensString, "\n")
		amountOfAliens, err := strconv.Atoi(amountOfAliensString)
		if err != nil {
			log.Fatal(err)
		} else {
			for index := 0; index < amountOfAliens; index++ {
				city := cityIndex[rand.Intn(len(cities))]
				aliens[index] = city
				writeLog("New Alien")
				writeLog(strconv.Itoa(index))
				writeLog(city.Name)
				writeLog("Total Aliens at City: " + strconv.Itoa(len(city.Aliens)))
				//TODO: Failing to add an alien to a city
				fmt.Print("Before: " + strconv.Itoa(len(city.Aliens)))
				city.Aliens = append(city.Aliens, index)
				fmt.Print("After: " + strconv.Itoa(len(city.Aliens)))
			}
			distributedAliens = true
		}
	}
}

func initializeLog() {
	_, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
}

func writeLog(message string) {
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	_, err = fmt.Fprintln(f, message)
	if err != nil {
		fmt.Println(err)
		f.Close()
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}
}
