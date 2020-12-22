package tokens

import (
	"fmt"
	"regexp"
	"strings"
)

func MakeLineTokens(container ContainerToken) {
	varRegex, _ := regexp.Compile("(<-)")

	if len(varRegex.FindStringIndex(container.Value)) != 0 {
		pos := varRegex.FindStringIndex(container.Value)
		variable := strings.TrimSpace(container.Value[:pos[0]])
		value := strings.TrimSpace(container.Value[pos[1]:])

		fmt.Println(VarToken{Id: container.Id, Variable: variable, Value: value})
	} else {
		fmt.Println(container)
	}
}