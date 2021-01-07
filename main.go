package main

import (
	"./supporting/parser"
	"./supporting/structs"
	"./supporting/tokens"
	"fmt"
)

func main() {
	// create empty array of tokens
	var tokenList []structs.Token

	// read lines of file to array
	lines := tokens.ReadFileLines("file.cnl")
	// split lines up into container tokens
	containerTokenList := tokens.MakeContainerTokens(lines)
	// iterate through container tokens, splitting into normal tokens
	for i := 0; i < len(containerTokenList); i++ {
		token := tokens.MakeLineTokens(containerTokenList[i])
		tokenList = append(tokenList, token)
	}

	// group together tokens in blocks
	tokenList = tokens.GroupBlocks(tokenList)

	fmt.Println(tokenList)

	// run parser
	// iterate through array of tokens
	for i := 0; i < len(tokenList); i++ {
		parser.Parse(tokenList[i], tokenList)
	}
}
