package tokens

import (
	"Container-lang/supporting/structs"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func MakeLineTokens(container structs.ContainerToken, containerPos int, containerTokenList []structs.ContainerToken) structs.Token {
	// compile regexes for checking if variable operation, function or creating block
	varRegex, _ := regexp.Compile("(<-)")
	funcRegex, _ := regexp.Compile("(PRINT\\s|EXECUTE\\s)")
	blockRegex, _ := regexp.Compile("BLOCK\\s")

	// check if variable or function token should be created
	if len(varRegex.FindStringIndex(container.Value)) != 0 { // create variable token
		// split container value based on <- position
		pos := varRegex.FindStringIndex(container.Value)
		variable := strings.TrimSpace(container.Value[:pos[0]])
		value := strings.TrimSpace(container.Value[pos[1]:])

		// return new token struct with only relevant variable field filled
		return structs.Token{Id: container.Id, VarToken: structs.VarToken{Variable: variable, Value: value}}
	} else if len(funcRegex.FindStringIndex(container.Value)) != 0 { // create function token
		// split container value by function name position
		pos := funcRegex.FindStringIndex(container.Value)
		function := strings.TrimSpace(container.Value[pos[0]:pos[1]])
		arguments := strings.TrimSpace(container.Value[pos[1]:])

		// return new token struct with only relevant function field filled
		return structs.Token{Id: container.Id, FunctionToken: structs.FunctionToken{Function: function, Arguments: arguments}}
	} else if len(blockRegex.FindStringIndex(container.Value)) != 0 && containerPos != -1 { // create block token
		// split container value based on "BLOCK" position and run error handling for entering non-int as argument
		pos := blockRegex.FindStringIndex(container.Value)
		arguments := strings.TrimSpace(container.Value[pos[1]:])
		argumentsInt, err := strconv.Atoi(arguments)
		if err != nil {
			log.Fatal("Runtime error: Container ID " + strconv.Itoa(container.Id) + ": Unable to convert '" + arguments + "' to integer")
		}

		// create array to append relevant containers to for returning
		var tokensToReturn []structs.Token

		// iterate through list of container tokens based on the position of block container and its argument
		for i := containerPos+1; i < containerPos+argumentsInt+1; i++ {
			// get create token to append to returning list, parsing in -1 as pos to prevent placing blocks in blocks
			tokensToReturn = append(tokensToReturn, MakeLineTokens(containerTokenList[i], -1, containerTokenList))
		}

		// return token with list of tokens as block attribute
		return structs.Token{Id: container.Id, Block: tokensToReturn}
	} else {
		log.Fatal("Runtime error: Container ID " + strconv.Itoa(container.Id) + ": Unrecognised request: " + container.Value)
	}

	return structs.Token{}
}