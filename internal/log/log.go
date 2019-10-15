package log

import (
	"fmt"
	. "github.com/cmwaters/alien_invasion/internal/general"
	"os"
)

func Initialize(logType string) {
	f, err := os.Create("log/" + logType + "_log.txt")
	Check(err)
	Write(logType, logType+" log opened")
	err = f.Close()
	Check(err)
}

func Write(logType string, message string) {
	f, err := os.OpenFile("log/"+logType+"_log.txt", os.O_APPEND|os.O_WRONLY, 0644)
	Check(err)
	_, err = fmt.Fprintln(f, message)
	Check(err)
	err = f.Close()
	Check(err)
}
