package parser

import (
	"../structs"
	"fmt"
)

func Parse(token structs.Token, tokenList []structs.Token) {
	// decide if function is being called or variable is being operated on
	if token.FunctionToken.Function == "" { // run variable stuff
		fmt.Println("variable")
	} else if token.VarToken.Variable == "" { // run function stuff
		fmt.Println("function")
	}
}