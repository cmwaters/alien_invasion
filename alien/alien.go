// Package to create the alien struct
package alien

import (
	"strconv"
)

type Alien struct {
	Id   int // must be unique
	City string
}

func (alien Alien) Name() string {
	return "Alien " + strconv.Itoa(alien.Id)
}
