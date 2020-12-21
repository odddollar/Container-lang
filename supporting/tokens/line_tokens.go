package tokens

import "strings"

func MakeLineTokens(fileLines []string) []string {
	var lineTokens []string

	for i := 0; i < len(fileLines); i++ {
		lineSplit := strings.Split(fileLines[i], "")

		start := 0

		for j := 0; j < len(lineSplit); j++ {
			if lineSplit[j] == "{" {
				start = j + 1
			}
			if lineSplit[j] == "}" {
				lineTokens = append(lineTokens, strings.Join(lineSplit[start:j], ""))
			}
		}
	}

	return lineTokens
}
