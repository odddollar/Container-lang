package parser

import (
	"Container-lang/supporting/structs"
	"github.com/Knetic/govaluate"
	"log"
	"strconv"
)

func createExpression(exp string, token structs.Token) interface{} {
	// create expression from exp argument
	expression, err := govaluate.NewEvaluableExpression(exp)
	if err != nil {
		log.Fatal("Runtime Error: Container ID " + strconv.Itoa(token.Id) + ": Unable to create expression from '" + token.VarToken.Value + "'")
	}

	// create dictionary to use as parameters
	params := make(map[string]interface{}, 64)
	for i := 0; i < len(variables); i++ {
		// add variable as number to dictionary
		params[variables[i].Name], _ = strconv.ParseFloat(variables[i].Value, 64)
	}

	// evaluate expression
	result, _ := expression.Evaluate(params)

	// check that result is returned to prevent entering strings
	if result == nil {
		log.Fatal("Runtime error: Container ID " + strconv.Itoa(token.Id) + ": Invalid numerical expression/number '" + token.VarToken.Value + "'")
	}

	return result
}