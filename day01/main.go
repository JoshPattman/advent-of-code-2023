package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	F "github.com/JoshPattman/functional"
)

// I know there are faster ways to do this, but this allows us to maintain compatability with pt2.
var conversionTablePt1 = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"0": 0,
}

// Tells us which words correspond to which numbers.
var conversionTablePt2 = map[string]int{
	"one":   1,
	"1":     1,
	"two":   2,
	"2":     2,
	"three": 3,
	"3":     3,
	"four":  4,
	"4":     4,
	"five":  5,
	"5":     5,
	"six":   6,
	"6":     6,
	"seven": 7,
	"7":     7,
	"eight": 8,
	"8":     8,
	"nine":  9,
	"9":     9,
	"zero":  0,
	"0":     0,
}

// This is what we will actually use for this run.
var conversionTable map[string]int

func main() {
	// Check if we should use word numbers too
	useLetterNumbers := flag.Bool("use-letter-numbers", false, "Use letter numbers too")
	flag.Parse()
	if *useLetterNumbers {
		conversionTable = conversionTablePt2
	} else {
		conversionTable = conversionTablePt1
	}

	// Load the input file.
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	inputFileData, err := io.ReadAll(inputFile)
	if err != nil {
		panic(err)
	}

	// Read all lines
	lines := strings.Split(string(inputFileData), "\n")
	// Calculate the number for each line
	digits := F.Map(lines, twoDigitLineNumber)
	// Sum all the numbers
	sum := F.Accumulate(digits, 0, func(a, b int) int { return a + b })
	// Print the output
	fmt.Println("Sum is", sum)
}

// Get the two digit number from a line
func twoDigitLineNumber(line string) int {
	firstDigit := ""
	lastDigit := ""
	for i := 0; i < len(line); i++ {
		if number, ok := startingNumber(line[i:]); ok {
			firstDigit = strconv.Itoa(number)
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		if number, ok := endingNumber(line[:i+1]); ok {
			lastDigit = strconv.Itoa(number)
			break
		}
	}
	s, _ := strconv.Atoi(firstDigit + lastDigit)
	return s
}

// Find the first number on the line
func startingNumber(s string) (int, bool) {
	for k, v := range conversionTable {
		if strings.HasPrefix(s, k) {
			return v, true
		}
	}
	return 0, false
}

// Find the last nubmer on a line
func endingNumber(s string) (int, bool) {
	for k, v := range conversionTable {
		if strings.HasSuffix(s, k) {
			return v, true
		}
	}
	return 0, false
}
