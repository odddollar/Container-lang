package functions

import (
	"../../structs"
	"fmt"
	"github.com/Knetic/govaluate"
	"log"
	"strconv"
)

func Print(text string, currentContainerID int, varList []structs.Variable) {
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

	// check that result was returned, otherwise log error
	if result == nil {
		log.Fatal("Runtime error: Container ID " + strconv.Itoa(currentContainerID) + ": Invalid numerical expression/variable '" + text + "'")
	}

	// print result
	fmt.Println(result)
}
