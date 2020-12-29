package functions

import (
	"fmt"
	"log"
	"strconv"
)

func Print(text string, currentContainerID int) {
	// attempt to convert argument to print to int, if unable log error
	if _, err := strconv.Atoi(text); err != nil {
		log.Fatal("Runtime error: Container ID " + strconv.Itoa(currentContainerID) + ": Invalid number '" + text + "'")
	} else {
		fmt.Println(text)
	}
}
