package parser

import (
	"../structs"
	"./functions"
	"fmt"
	"github.com/Knetic/govaluate"
	"log"
	"strconv"
)

// create array of variables
var variables []structs.Variable

func Parse(token structs.Token, tokenList []structs.Token) {
	// decide if function is being called or variable is being operated on
	if token.FunctionToken.Function == "" { // run variable stuff
		// check if variable is in variable list, if not add to list with empty values
		if !checkVarExists(token.VarToken.Variable, variables) {
			variables = append(variables, structs.Variable{Name: token.VarToken.Variable})
		}

		// create expression from token value
		expression, err := govaluate.NewEvaluableExpression(token.VarToken.Value)
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

		// assign value to variable in variable array
		varPos := getVarPosByName(token.VarToken.Variable, variables)
		variables[varPos].Value = fmt.Sprintf("%v", result)

	} else if token.VarToken.Variable == "" { // run function stuff
		if token.FunctionToken.Function == "PRINT" { // run print function
			functions.Print(token.FunctionToken.Arguments, token.Id)

		} else if token.FunctionToken.Function == "EXECUTE" { // run execute stuff
			// get id of container to execute
			idToExecute, err := strconv.Atoi(token.FunctionToken.Arguments)
			if err != nil {
				log.Fatal("Runtime error: Container ID " + strconv.Itoa(token.Id) + ": Non-numerical ID supplied")
			}

			// return token after finding it in list
			executedToken := getContainerById(idToExecute, tokenList, token.Id)

			// recursively call parser function with new token
			Parse(executedToken, tokenList)
		}
	}

	fmt.Println(variables)
}