package main

import (
	"./supporting/tokens"
	"fmt"
)

func main() {
	lines := tokens.ReadFileLines("file.seu")
	fmt.Println(tokens.MakeLineTokens(lines))
}
