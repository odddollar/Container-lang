package tokens

import (
	"../structs"
	"fmt"
)

func GroupBlocks(tokenList []structs.Token) []structs.Token {
	for i := 0; i < len(tokenList); i++ {
		fmt.Println(tokenList[i])
	}

	return tokenList
}