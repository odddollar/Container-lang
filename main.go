package main

import (
	"./supporting/tokens"
	"fmt"
)

func main() {
	lines := tokens.ReadFileLines("file.cnl")
	containerTokenList := tokens.MakeContainerTokens(lines)
	for i := 0; i < len(containerTokenList); i++ {
		varToken, functionToken := tokens.MakeLineTokens(containerTokenList[i])
		if functionToken.Id == -1 {
			fmt.Println(varToken)
		} else if varToken.Id == -1 {
			fmt.Println(functionToken)
		}
	}
}
