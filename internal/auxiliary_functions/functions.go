package auxiliary_functions

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func NextWordIsADirection(sentence string, index int) bool {
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

func GetFileName() string {
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
