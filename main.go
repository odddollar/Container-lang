package main

import (
	"Container-lang/supporting/parser"
	"Container-lang/supporting/structs"
	"Container-lang/supporting/tokens"
	"fmt"
	"github.com/akamensky/argparse"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	// get start time
	start := time.Now()

	// create variables to allow for argparser
	var argparser *argparse.Parser
	var inputFile *string
	var debug *bool

	// create lines array to read input file
	var lines []string

	// create argument parser
	argparser = argparse.NewParser("Container-lang", "Interpreter for the Container-lang programming language")
	inputFile = argparser.String("f", "file", &argparse.Options{Required: true, Help: "The path to the input .cnl file"})
	debug = argparser.Flag("d", "debug", &argparse.Options{Required: false, Help: "Print token stream and program completion time"})

	// run argument parser
	if err := argparser.Parse(os.Args); err != nil {
		log.Fatal(argparser.Usage(err))
	}

	// check that file uses .cnl extension 'cause why not
	if !strings.Contains(*inputFile, ".cnl") {
		log.Fatal("File error: Ensure file uses .cnl extension")
	}

	// create empty array of tokens
	var tokenList []structs.Token

	// read lines of file to array
	lines = tokens.ReadFileLines(*inputFile)

	// split lines up into container tokens
	containerTokenList := tokens.MakeContainerTokens(lines)

	// iterate through container tokens, splitting into normal tokens
	for i := 0; i < len(containerTokenList); i++ {
		token := tokens.MakeLineTokens(containerTokenList[i], i, containerTokenList)
		tokenList = append(tokenList, token)

		// check if most recent token was a block
		if len(token.Block) != 0 {
			i += tokens.GetTotalNumberTokensInBlock(token.Block)
		}
	}

	// print token list if debug flag given
	if *debug {
		fmt.Println(tokenList)
	}

	// run parser
	// iterate through array of tokens
	for i := 0; i < len(tokenList); i++ {
		parser.Parse(tokenList[i], &tokenList)
	}

	// print completion time if debug flag given
	if *debug {
		fmt.Printf("Finished in: %v", time.Since(start))
	}
}
