package tokens

import (
	"regexp"
	"strconv"
	"strings"
)

func MakeLineTokens(fileLines []string) []ContainerToken {
	var lineTokens []ContainerToken

	idRegexFromLIne, _ := regexp.Compile("{\\s*\\d*\\s*\\|")
	idRegex, _ := regexp.Compile("\\d+")

	for i := 0; i < len(fileLines); i++ {
		lineSplit := strings.Split(fileLines[i], "")

		start := 0

		for j := 0; j < len(lineSplit); j++ {
			if lineSplit[j] == "{" {
				start = j
			}
			if lineSplit[j] == "}" {
				container := strings.Join(lineSplit[start:j+1], "")

				idPartPos := idRegexFromLIne.FindStringIndex(container)

				id, _ := strconv.Atoi(idRegex.FindString(container[idPartPos[0]:idPartPos[1]]))
				line := strings.TrimSpace(strings.ReplaceAll(container[idPartPos[1]:], "}", ""))

				lineTokens = append(lineTokens, ContainerToken{Id: id, Value: line})
			}
		}
	}

	return lineTokens
}
