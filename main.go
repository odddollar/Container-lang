package main

import (
	"fmt"
	"strings"
)

func main() {
	str := strings.Split(`(CREATE x)(SET x=1)`, "")

	var lines []string

	start := 0

	for i := 0; i < len(str); i++ {
		if str[i] == "(" {
			start = i + 1
		}
		if str[i] == ")" {
			end := i
			lines = append(lines, strings.Join(str[start:end], ""))
		}
	}

	fmt.Println(lines)
}
