package parser

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"log"
	"strconv"
)

func print_(text string, currentContainerID int) {
	// create govaluate expression from parsed in print parameters
	expression, err := govaluate.NewEvaluableExpression(text)
	if err != nil {
		log.Fatal("Runtime Error: Container ID " + strconv.Itoa(currentContainerID) + ": Unable to create expression from '" + text + "'")
	}

	// create dictionary of variables and their values
	params := make(map[string]interface{}, 64)
	for i := 0; i < len(variables); i++ {
		// add variable as number to dictionary
		params[variables[i].Name], _ = strconv.ParseFloat(variables[i].Value, 64)
	}

	// evaluate expression
	result, _ := expression.Evaluate(params)

	// check that result was returned, otherwise log error
	if result == nil {
		log.Fatal("Runtime error: Container ID " + strconv.Itoa(currentContainerID) + ": Invalid numerical expression/variable '" + text + "'")
	}

	// print result
	fmt.Println(result)
}
