package parser

import (
	"../structs"
	"./functions"
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

		// assign value to variable

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
}