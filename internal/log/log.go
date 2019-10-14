package log

import (
	"fmt"
	"os"
)

func Initialize(logType string) {
	f, err := os.Create("log/" + logType + "_log.txt")
	if err != nil {
		fmt.Println(err)
	}
	Write(logType, logType+" log opened")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func Write(logType string, message string) {
	f, err := os.OpenFile("log/"+logType+"_log.txt", os.O_APPEND|os.O_WRONLY, 0644)
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
