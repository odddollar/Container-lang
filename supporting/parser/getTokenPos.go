package parser

import (
	"Container-lang/supporting/structs"
)

func getTokenPos(tokenID int, tokenList []structs.Token) int {
	for i := 0; i < len(tokenList); i++ {
		if tokenList[i].Id == tokenID {
			return i
		}
	}

	return 0
}
