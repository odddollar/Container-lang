package tokens

import (
	"regexp"
	"strings"
)

func MakeLineTokens(container ContainerToken) (VarToken, FunctionToken) {
	// compile regexes for checking if variable operation or function
	varRegex, _ := regexp.Compile("(<-)")
	funcRegex, _ := regexp.Compile("(PRINT)")

	// check if variable or function token should be created
	if len(varRegex.FindStringIndex(container.Value)) != 0 {
		// split container value based on <- position
		pos := varRegex.FindStringIndex(container.Value)
		variable := strings.TrimSpace(container.Value[:pos[0]])
		value := strings.TrimSpace(container.Value[pos[1]:])

		return VarToken{Id: container.Id, Variable: variable, Value: value}, FunctionToken{Id: -1}
	} else if len(funcRegex.FindStringIndex(container.Value)) != 0 {
		// split container value by function name position
		pos := funcRegex.FindStringIndex(container.Value)
		function := strings.TrimSpace(container.Value[pos[0]:pos[1]])
		arguments := strings.TrimSpace(container.Value[pos[1]:])

		return VarToken{Id: -1}, FunctionToken{Id: container.Id, Function: function, Arguments: arguments}
	}

	return VarToken{}, FunctionToken{}
}