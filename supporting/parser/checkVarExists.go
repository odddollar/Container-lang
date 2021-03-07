package parser

import (
	"Container-lang/supporting/structs"
)

func checkVarExists(varName string, varList *[]structs.Variable) bool {
	// iterate through list of variables
	for i := 0; i < len(*varList); i++ {
		// if variable is found return true
		if (*varList)[i].Name == varName {
			return true
		}
	}

	// if variable isn't found return false
	return false
}
