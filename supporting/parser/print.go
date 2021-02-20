package parser

import (
	"Container-lang/supporting/structs"
	"fmt"
)

func print_(text string, token structs.Token) {
	result := createExpression(text, token)

	// print result
	fmt.Println(result)
}
