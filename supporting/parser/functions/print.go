package functions

import (
	"../../structs"
	"fmt"
	"github.com/Knetic/govaluate"
	"log"
	"regexp"
	"strconv"
)

func Print(text string, currentContainerID int, varList []structs.Variable) {
	// create regex to search for only letters
	findLetters, _ := regexp.Compile("[a-zA-Z]*")

	// get locations of any letters in the parsed text-to-print
	letterLocations := findLetters.FindStringIndex(text)

	// check if text is a variable and if letters are present to prevent printing of raw strings but allow printing of expressions and variables
	if (letterLocations[0] != 0 || letterLocations[1] != 0) && checkIfVar(text, varList) == false {
		log.Fatal("Runtime error: Container ID " + strconv.Itoa(currentContainerID) + ": Invalid numerical expression/variable '" + text + "'")
	} else {
		// create govaluate expression from parsed in print parameters
		expression, err := govaluate.NewEvaluableExpression(text)
		if err != nil {
			log.Fatal("Runtime Error: Container ID " + strconv.Itoa(currentContainerID) + ": Unable to create expression from '" + text + "'")
		}

		// create dictionary of variables and their values
		params := make(map[string]interface{}, 64)
		for i := 0; i < len(varList); i++ {
			// add variable as number to dictionary
			params[varList[i].Name], _ = strconv.ParseFloat(varList[i].Value, 64)
		}

		// evaluate expression
		result, _ := expression.Evaluate(params)

		// print result
		fmt.Println(result)
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