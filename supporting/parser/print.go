package parser

import (
	"Container-lang/supporting/structs"
	"fmt"
)

func print_(text string, token structs.Token, newLine bool) {
	result := createExpression(text, token)

	// print result
	if newLine == true {
		fmt.Println(result)
	} else {
		fmt.Print(result)
	}
}
