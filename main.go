package main

import (
	"./supporting/tokens"
	"fmt"
)

func main() {
	var tokenList []tokens.Token

	lines := tokens.ReadFileLines("file.cnl")
	containerTokenList := tokens.MakeContainerTokens(lines)
	for i := 0; i < len(containerTokenList); i++ {
		token := tokens.MakeLineTokens(containerTokenList[i])
		tokenList = append(tokenList, token)

		fmt.Println(token)
	}

	fmt.Println(tokenList)
}
