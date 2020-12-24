package main

import (
	"./supporting/tokens"
	"fmt"
)

func main() {
	lines := tokens.ReadFileLines("file.cnl")
	containerTokenList := tokens.MakeContainerTokens(lines)
	for i := 0; i < len(containerTokenList); i++ {
		token := tokens.MakeLineTokens(containerTokenList[i])
		fmt.Println(token)
	}
}
