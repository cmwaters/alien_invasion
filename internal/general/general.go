package general

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

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
