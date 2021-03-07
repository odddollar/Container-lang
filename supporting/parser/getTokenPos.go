package parser

import (
	"Container-lang/supporting/structs"
)

func getTokenPos(tokenID int, tokenList *[]structs.Token) int {
	// deference back to slice
	tokens := *tokenList

	// iterate through tokens
	for i := 0; i < len(*tokenList); i++ {
		if tokens[i].Id == tokenID {
			return i
		}
	}

	return 0
}
