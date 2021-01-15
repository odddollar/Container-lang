package parser

import (
	"Container-lang/supporting/structs"
	"log"
	"strconv"
)

func getContainerById(id int, tokenList []structs.Token, currentTokenId int) structs.Token {
	// iterate through tokens in list
	for i := 0; i < len(tokenList); i++ {
		// check if current token matches requested id
		if tokenList[i].Id == id {
			return tokenList[i]
		}
	}

	// return fatal error if no container with requested id is found
	log.Fatal("Runtime error: Container ID "  + strconv.Itoa(currentTokenId) + ": No container with ID " + strconv.Itoa(id) + " found")
	return structs.Token{}
}
