package parser

import (
	"../structs"
	"./functions"
	"fmt"
)

func Parse(token structs.Token, tokenList []structs.Token) {
	// decide if function is being called or variable is being operated on
	if token.FunctionToken.Function == "" { // run variable stuff
		fmt.Println("variable")
	} else if token.VarToken.Variable == "" { // run function stuff
		if token.FunctionToken.Function == "PRINT" { // run print function
			functions.Print(token.FunctionToken.Arguments)
		} else if token.FunctionToken.Function == "EXECUTE" { // run execute stuff
			fmt.Println("execute function")
		}
	}
}