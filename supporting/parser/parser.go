package parser

import (
	"../structs"
	"./functions"
	"fmt"
	"strconv"
)

func Parse(token structs.Token, tokenList []structs.Token) {
	// decide if function is being called or variable is being operated on
	if token.FunctionToken.Function == "" { // run variable stuff
		fmt.Println("variable")
	} else if token.VarToken.Variable == "" { // run function stuff
		if token.FunctionToken.Function == "PRINT" { // run print function
			functions.Print(token.FunctionToken.Arguments, token.Id)
		} else if token.FunctionToken.Function == "EXECUTE" { // run execute stuff
			// get id of container to execute
			idToExecute, _ := strconv.Atoi(token.FunctionToken.Arguments)

			// return token after finding it in list
			executedToken := getContainerById(idToExecute, tokenList, token.Id)

			// recursively call parser function with new token
			Parse(executedToken, tokenList)
		}
	}
}