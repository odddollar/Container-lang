package tokens

import (
	"../structs"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func MakeContainerTokens(fileLines []string) []structs.ContainerToken {
	// take in array of file lines, create array of pre-processed container tokens
	var lineTokens []structs.ContainerToken

	// create regexes for finding id in code container
	idRegexFromLIne, _ := regexp.Compile("{\\s*\\d*\\s*\\|")
	idRegex, _ := regexp.Compile("\\d+")

	// iterate through lines in input file array
	for i := 0; i < len(fileLines); i++ {
		// split line into characters
		lineSplit := strings.Split(fileLines[i], "")

		start := -1

		// iterate through characters in current line
		for j := 0; j < len(lineSplit); j++ {
			// run when end of container found
			if lineSplit[j] == "}" && start != -1{
				// separate current container into string
				container := strings.Join(lineSplit[start:j+1], "")

				// get position of id part of container using regex
				idPartPos := idRegexFromLIne.FindStringIndex(container)

				// check that an id has been provided
				if len(idPartPos) == 0 {
					log.Fatal("Syntax error: Line " + strconv.Itoa(i+1) + ": No valid container ID found")
				}

				// find id within first part of container and convert to int
				id, _ := strconv.Atoi(idRegex.FindString(container[idPartPos[0]:idPartPos[1]]))

				// get second part of container string and remove bad characters for value in container
				line := strings.TrimSpace(strings.ReplaceAll(container[idPartPos[1]:], "}", ""))

				// append container token to list
				lineTokens = append(lineTokens, structs.ContainerToken{Id: id, Value: line})

				// reset start to -1 to search for missing {
				start = -1
			} else if lineSplit[j] == "}" && start == -1 { // check for missing {
				log.Fatal("Syntax error: Line " + strconv.Itoa(i+1) + ": Found } with no {")
			} else if (start != -1 && j == len(lineSplit)-1) || (start != -1 && lineSplit[j] == "{") { // check for missing }
				log.Fatal("Syntax error: Line " + strconv.Itoa(i+1) + ": Found { with no }")
			} else if lineSplit[j] == "{" { // set start to current position when opening of container found. must run after } search in order to allow missing } syntax error
				start = j
			}
		}
	}

	return lineTokens
}
