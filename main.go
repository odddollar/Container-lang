package main

import (
	"./supporting/tokens"
)

func main() {
	lines := tokens.ReadFileLines("file.cnl")
	containerTokenList := tokens.MakeContainerTokens(lines)
	for i := 0; i < len(containerTokenList); i++ {
		tokens.MakeLineTokens(containerTokenList[i])
	}
}
