package parser

import (
	"Container-lang/supporting/structs"
)

func getVarPosByName(varName string, varList *[]structs.Variable) int {
	// iterate through list of variables
	for i := 0; i < len(*varList); i++ {
		// if names match, return the position of variable in array
		if (*varList)[i].Name == varName {
			return i
		}
	}

	return 0
}
