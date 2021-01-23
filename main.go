package main

import (
	"Container-lang/supporting/parser"
	"Container-lang/supporting/structs"
	"Container-lang/supporting/tokens"
	"fmt"
	"github.com/akamensky/argparse"
	"log"
	"os"
)

func main() {
	var argparser *argparse.Parser
	var inputFile *string
	var lines []string

	// change for testing in development and release
	developing := false

	// run argparse stuff if not developing
	if developing == false {
		// create argument parser
		argparser = argparse.NewParser("Container-lang", "Interpreter for the Container-lang programming language")
		inputFile = argparser.String("f", "file", &argparse.Options{Required: true})

		// run argument parser
		if err := argparser.Parse(os.Args); err != nil {
			log.Fatal(argparser.Usage(err))
		}
	}

	// create empty array of tokens
	var tokenList []structs.Token

	// read lines of file to array, checking whether argument parsed in or use test file
	if developing == true {
		lines = tokens.ReadFileLines("file.cnl")
	} else {
		lines = tokens.ReadFileLines(*inputFile)
	}

	// split lines up into container tokens
	containerTokenList := tokens.MakeContainerTokens(lines)

	// iterate through container tokens, splitting into normal tokens
	for i := 0; i < len(containerTokenList); i++ {
		token := tokens.MakeLineTokens(containerTokenList[i], i, containerTokenList)
		tokenList = append(tokenList, token)

		// check if most recent token was a block
		if len(token.Block) != 0 {
			i += len(token.Block)
		}
	}

	fmt.Println(tokenList)

	// run parser
	// iterate through array of tokens
	for i := 0; i < len(tokenList); i++ {
		parser.Parse(tokenList[i], tokenList)
	}
}
