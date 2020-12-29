package parser

import (
	"../structs"
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

	log.Fatal("Runtime error: Container ID "  + strconv.Itoa(currentTokenId) + ": No token with ID " + strconv.Itoa(id) + " found")
	return structs.Token{}
}
