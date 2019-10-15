# alien_invasion

### Summary

This project simulates an alien invasion by reading an input file that contains a list of cities
and their neighbor cities i.e. `Auckland north=Whangarei south=Hamilton east=Coromandel` 
It then populates these cities with a chosen number of aliens. At each step in the simulation, 
aliens move from one city to one of its random neighboring cities. After this movement step, if two or
more aliens are in the same city, they fight. This destroys the aliens in this city as well as the city 
itself (such that aliens can no longer move across this city to another). The simulation runs for 10000 
steps or until all aliens are killed. The simulation outputs a txt file of the same format with the new list
of cities that remain following the alien invasion.

### Build

In the terminal/console from the ROOT directory run:

`go build`

To install to your /bin directory in your workspace run:

`go install ./...`

### Run

In the terminal/console from the ROOT directory run:

`.\alien_invasion`

The console will ask you to pick a file from the map directory with which to run the simulation with. 
You can load .txt files of the right format or generate a .txt file using the generate package. 
The simulation will then ask for the number of aliens which will be randomly distributed between cities.
Upon finishing of the simulation a corresponding output.txt file will be written in the output folder


### Testing

Several tests have been used to check that the functions perform as expected. Any adjustments should also
make sure that the initial functions still pass the tests. Upon building of the package you can run 

`go test`

to test the functionality of the packages. 

### Directory Structure

The architecture of the directory is as follows:

```
main.go // the main package that the simulation runs from
main_test.go // the testing of the main package
doc.go // to support the documentation
-> alien // supports the alien struct and related funcs
    alien.go
-> city // supports the city struct and related funcs
    city.go
    city_test.go
-> general // supports all other packages - no dependencies
    general.go
-> generate // generates city maps and output files
    generate.go
    generate_test.go
-> log // used for logging for debugging and errors
    log.go
    *debug_log.txt*
    *error_log.txt*
-> maps // folder where maps are kept
    *test_map.txt*
-> output // folder where the output to each simulation is kept
    *output.txt*
```

### Documentation

For a more extensive documentation of the packages, use Godoc:

`godoc -http=:6060`

and then go to https://localhost.com/6060/pkg and find the 
alien_invasion package in the directory