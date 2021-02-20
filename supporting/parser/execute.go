package parser

import "Container-lang/supporting/structs"

func execute(idToExecute int, token structs.Token, tokenList []structs.Token) {
	// return token after finding it in list
	executedToken := getContainerById(idToExecute, tokenList, token.Id)

	// check if executing block or normal container
	if len(executedToken.Block) == 0 {
		// recursively call parser function with new token
		Parse(executedToken, tokenList)
	} else {
		// iterate through tokens in block
		for i := 0; i < len(executedToken.Block); i++ {
			// recursively call parser function with current token and list of tokens in block
			Parse(executedToken.Block[i], executedToken.Block)
		}
	}
}
