/*
Package alien_invasion simulates an alien invasion to a collection of cities

Summary

This project simulates an alien invasion by reading an input file that contains a list of cities and their neighbor cities
i.e. Auckland north=Whangarei south=Hamilton east=Coromandel It then populates these cities with a chosen number of aliens.
At each step in the simulation, aliens move from one city to one of its random neighboring cities. After this movement
step, if two or more aliens are in the same city, they fight. This destroys the aliens in this city as well as the city
itself (such that aliens can no longer move across this city to another). The simulation runs for 10000 steps or until
all aliens are killed. The simulation outputs a txt file of the same format with the new list of cities that remain
following the alien invasion.


Constants and Structs

This is used to keep a running collection of the active cities and the aliens that they have in each as seen in the City struct.

	cities map[string]*c.City

The City struct:

	type City struct {
		Name   string
		North  *City
		East   *City
		West   *City
		South  *City
		Aliens map[int]*a.Alien
	}

The aliens map allows us to loop through all the aliens and the cities in as opposed to loop through the cities

	aliens map[int]*a.Alien

The Alien struct

	type Alien struct {
		Id   int
		City string
	}

This constant defined the size (in this case 3 x 3) of the generated map

	WorldSize = 3

Functions


main() :
the main function begins when we run the executable file. This initialises the maps: cities and aliens, it then
implements other functions to retrieve the information from the input file and the amount of aliens. It then calls the inject_aliens()
function to disperse the aliens randomly over the map. The main simulation is then run for 10000 iterations or when the aliens
are destroyed. At each iteration the step() function is called. Upon exit, this function produces an output file

step() :
Steps through one iteration of the simulation. This consists of moving each Aliens to a randomly selected
neighbor city to the one it currently resides in and then at the close checking if any of the cities contain
two or more aliens in which both the city and the aliens inhibiting it are destroyed and removed from the map

processFile() :
Takes a txt file (as a string) and iterates over each line in search for cities to add to the cities map that it returns

injectAliens() :
This function retrieves the amount of aliens the user wants and adds them to a random city

moveAlien(alien a.Alien)
This function moves an alien from its current city to a randomly chosen neighboring city

evaluateCity(city *c.City)
This function checks how many aliens in a city. If greater than or equal 2, the function proceeds to output the results
to the console and remove both the aliens and the city from their respective maps.




*/
package main
