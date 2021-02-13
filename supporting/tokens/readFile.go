package tokens

import (
	"bufio"
	"log"
	"os"
)

func ReadFileLines(dir string) []string {
	// open file
	file, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}

	var lines []string

	// create file scanner
	scanner := bufio.NewScanner(file)

	// iterate through file lines, appending to array
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
