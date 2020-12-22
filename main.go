package main

import (
	"./supporting/tokens"
	"fmt"
)

func main() {
	lines := tokens.ReadFileLines("file.cnl")
	fmt.Println(tokens.MakeContainerTokens(lines))
}
