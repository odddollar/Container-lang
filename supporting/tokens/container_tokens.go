package tokens

import (
	"regexp"
	"strconv"
	"strings"
)

func MakeLineTokens(fileLines []string) []ContainerToken {
	// take in array of file lines, create array of pre-processed container tokens
	var lineTokens []ContainerToken

	// create regexes for finding id in code container
	idRegexFromLIne, _ := regexp.Compile("{\\s*\\d*\\s*\\|")
	idRegex, _ := regexp.Compile("\\d+")

	// iterate through lines in input file array
	for i := 0; i < len(fileLines); i++ {
		// split line into characters
		lineSplit := strings.Split(fileLines[i], "")

		start := 0

		// iterate through characters in current line
		for j := 0; j < len(lineSplit); j++ {
			// set start to current position when opening of container found
			if lineSplit[j] == "{" {
				start = j
			}
			// run when end of container found
			if lineSplit[j] == "}" {
				// separate current container into string
				container := strings.Join(lineSplit[start:j+1], "")

				// get position of id part of container using regex
				idPartPos := idRegexFromLIne.FindStringIndex(container)

				// find id within first part of container and convert to int
				id, _ := strconv.Atoi(idRegex.FindString(container[idPartPos[0]:idPartPos[1]]))

				// get second part of container string and remove bad characters for value in container
				line := strings.TrimSpace(strings.ReplaceAll(container[idPartPos[1]:], "}", ""))

				// append container token to list
				lineTokens = append(lineTokens, ContainerToken{Id: id, Value: line})
			}
		}
	}

	return lineTokens
}
