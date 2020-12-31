package functions

import (
	"../../structs"
	"fmt"
	"log"
	"strconv"
)

func Print(text string, currentContainerID int, varList []structs.Variable) {
	// attempt to convert argument to print to int and not attempting to print variable, if unable log error
	if _, err := strconv.Atoi(text); err != nil && !checkIfVar(text, varList) {
		log.Fatal("Runtime error: Container ID " + strconv.Itoa(currentContainerID) + ": Invalid number/variable '" + text + "'")
	} else {
		// check if printing variable
		if checkIfVar(text, varList) {
			// get variable position in varList then print value
			varPos := getVarPosByName(text, varList)
			fmt.Println(varList[varPos].Value)
		} else {
			// print plain text
			fmt.Println(text)
		}
	}
}

func checkIfVar(varName string, varList []structs.Variable) bool {
	// iterate through list of variables
	for i := 0; i < len(varList); i++ {
		// if variable is found return true
		if varList[i].Name == varName {
			return true
		}
	}

	// if variable isn't found return false
	return false
}

func getVarPosByName(varName string, varList []structs.Variable) int {
	// iterate through list of variables
	for i := 0; i < len(varList); i++ {
		// if names match, return the position of variable in array
		if varList[i].Name == varName {
			return i
		}
	}

	return 0
}