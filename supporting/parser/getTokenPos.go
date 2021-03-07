package parser

import (
	"Container-lang/supporting/structs"
)

func getTokenPos(tokenID int, tokenList *[]structs.Token) int {
	// iterate through tokens
	for i := 0; i < len(*tokenList); i++ {
		if (*tokenList)[i].Id == tokenID {
			return i
		}
	}

	return 0
}
