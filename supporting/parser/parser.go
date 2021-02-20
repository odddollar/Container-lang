package parser

import (
	"Container-lang/supporting/structs"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// create array of variables
var variables []structs.Variable

func Parse(token structs.Token, tokenList []structs.Token) {
	// decide if function is being called or variable is being operated on
	if token.FunctionToken.Function == "" && token.VarToken.Variable != "" { // run variable stuff
		// check if variable is in variable list, if not add to list with empty values
		if !checkVarExists(token.VarToken.Variable, variables) {
			variables = append(variables, structs.Variable{Name: token.VarToken.Variable})
		}

		// assign value to variable in variable array
		varPos := getVarPosByName(token.VarToken.Variable, variables)
		variables[varPos].Value = fmt.Sprintf("%v", createExpression(token.VarToken.Value, token))

	} else if token.VarToken.Variable == "" && token.FunctionToken.Function != "" { // run function stuff
		if token.FunctionToken.Function == "PRINT" { // run print function
			print_(token.FunctionToken.Arguments, token.Id)

		} else if token.FunctionToken.Function == "REPEAT" {
			// split arguments
			args := strings.Split(token.FunctionToken.Arguments, ",")
			if len(args) != 2 {
				log.Fatal("Runtime error: Container ID " + strconv.Itoa(token.Id) + ": Required 2 arguments, " + strconv.Itoa(len(args)) + " provided")
			}

			// convert args to integers
			containerToRepeat, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal("Runtime error: Container ID " + strconv.Itoa(token.Id) + ": Non-numerical container ID supplied in argument 1")
			}

			// prevent creation of infinite loop with repeats calling themselves
			if containerToRepeat == token.Id {
				log.Fatal("Fatal error: Container ID " + strconv.Itoa(token.Id) + ": Repeat calling itself, infinite loop prevented")
			}

			// create expression from second argument and evaluate to allow for maths and variables
			repetitions := int(createExpression(args[1], token).(float64))

			// get positions of tokens in token array
			executePos := getTokenPos(token.Id, tokenList)
			toExecutePos := getTokenPos(containerToRepeat, tokenList)

			// only allow executing of container after first use
			if executePos < toExecutePos {
				log.Fatal("Runtime error: Container ID " + strconv.Itoa(token.Id) + ": Attempting to execute container prior to its definition")
			}

			// return token after finding it in list
			executedToken := getContainerById(containerToRepeat, tokenList, token.Id)

			// repeat number of repetitions
			for j := 0; j < repetitions; j++ {
				// check if executing block or normal container
				if len(executedToken.Block) == 0 {
					// recursively call parser function with new token
					Parse(executedToken, tokenList)
				} else {
					// iterate through tokens in block
					for i := 0; i < len(executedToken.Block); i++ {
						// recursively call parser function with current token and list of tokens in block
						Parse(executedToken.Block[i], executedToken.Block)
					}
				}
			}

		} else if token.FunctionToken.Function == "EXECUTE" { // run execute stuff
			// get id of container to execute
			idToExecute, err := strconv.Atoi(token.FunctionToken.Arguments)
			if err != nil {
				log.Fatal("Runtime error: Container ID " + strconv.Itoa(token.Id) + ": Non-numerical ID supplied")
			}

			// get positions of tokens in token array
			executePos := getTokenPos(token.Id, tokenList)
			toExecutePos := getTokenPos(idToExecute, tokenList)

			// only allow executing of container after first use
			if executePos < toExecutePos {
				log.Fatal("Runtime error: Container ID " + strconv.Itoa(token.Id) + ": Attempting to execute container prior to its definition")
			}

			// return token after finding it in list
			executedToken := getContainerById(idToExecute, tokenList, token.Id)

			// check if executing block or normal container
			if len(executedToken.Block) == 0 {
				// recursively call parser function with new token
				Parse(executedToken, tokenList)
			} else {
				// iterate through tokens in block
				for i := 0; i < len(executedToken.Block); i++ {
					// recursively call parser function with current token and list of tokens in block
					Parse(executedToken.Block[i], executedToken.Block)
				}
			}
		}
	}
}
