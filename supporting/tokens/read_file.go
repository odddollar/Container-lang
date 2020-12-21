package tokens

import (
	"bufio"
	"log"
	"os"
)

func ReadFileLines(dir string) []string {
	file, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}