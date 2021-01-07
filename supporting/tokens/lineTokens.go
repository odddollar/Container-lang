package tokens

import (
	"../structs"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func MakeLineTokens(container structs.ContainerToken) structs.Token {
	// compile regexes for checking if variable operation or function
	varRegex, _ := regexp.Compile("(<-)")
	funcRegex, _ := regexp.Compile("(PRINT\\s|EXECUTE\\s|BLOCK\\s)")

	// check if variable or function token should be created
	if len(varRegex.FindStringIndex(container.Value)) != 0 {
		// split container value based on <- position
		pos := varRegex.FindStringIndex(container.Value)
		variable := strings.TrimSpace(container.Value[:pos[0]])
		value := strings.TrimSpace(container.Value[pos[1]:])

		// return new token struct with only relevant variable field filled
		return structs.Token{Id: container.Id, VarToken: structs.VarToken{Variable: variable, Value: value}}
	} else if len(funcRegex.FindStringIndex(container.Value)) != 0 {
		// split container value by function name position
		pos := funcRegex.FindStringIndex(container.Value)
		function := strings.TrimSpace(container.Value[pos[0]:pos[1]])
		arguments := strings.TrimSpace(container.Value[pos[1]:])

		// return new token struct with only relevant function field filled
		return structs.Token{Id: container.Id, FunctionToken: structs.FunctionToken{Function: function, Arguments: arguments}}
	} else {
		log.Fatal("Runtime error: Container ID " + strconv.Itoa(container.Id) + ": Unrecognised request: " + container.Value)
	}

	return structs.Token{}
}