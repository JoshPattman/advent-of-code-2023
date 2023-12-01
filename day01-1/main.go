package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	F "github.com/JoshPattman/functional"
)

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	inputFileData, err := io.ReadAll(inputFile)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputFileData), "\n")
	digits := F.Map(lines, twoDigitLineNumber)
	sum := F.Accumulate(digits, 0, func(a, b int) int { return a + b })
	fmt.Println("Sum is", sum)
}

func twoDigitLineNumber(line string) int {
	firstDigit := ""
	lastDigit := ""
	for _, char := range line {
		if isDigit(char) {
			if firstDigit == "" {
				firstDigit = string(char)
			}
			lastDigit = string(char)
		}
	}
	s, _ := strconv.Atoi(firstDigit + lastDigit)
	return s
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
