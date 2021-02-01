package tokens

import (
	"Container-lang/supporting/structs"
)

func GetTotalNumberTokensInBlock(tokens []structs.Token) int {
	var total int

	// iterate through input tokens
	for i := 0; i < len(tokens); i++ {
		// check if the current token holds other tokens
		if len(tokens[i].Block) == 0 { // no other tokens present, add one
			total += 1
		} else { // other tokens present, recursively call function and add to total
			total += GetTotalNumberTokensInBlock(tokens[i].Block) + 1
		}
	}

	return total
}