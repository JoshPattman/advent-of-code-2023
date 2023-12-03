package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// Tracks a position on the schematic
type pos struct {
	x, y int
}

// Decides if a rune is a number
func isNum(r rune) bool {
	return r >= '0' && r <= '9'
}

// Returns a slice of unique elements from a slice
func unique[T comparable](xs []T) []T {
	seen := make(map[T]bool)
	out := make([]T, 0)
	for _, x := range xs {
		if !seen[x] {
			out = append(out, x)
			seen[x] = true
		}
	}
	return out
}

// Function to check (bounds checking included) if a position is a symbol
func isPossibleGear(matrix [][]rune, y, x int) bool {
	if x < 0 || y < 0 || y >= len(matrix) || x >= len(matrix[y]) {
		return false
	}
	return !isNum(matrix[y][x]) && matrix[y][x] != '.'
}

// Main entrypoint
func main() {
	// Read the input into a matrix
	lines := strings.Split(input, "\n")
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}

	// some stuff for storing what we find
	allParts := make(map[pos]int)         // pos (start of part) -> part number
	allGears := make(map[pos]map[pos]int) // pos (gear) -> (pos (start of part) -> part number)

	// Loop over every line normally
	for y := 0; y < len(matrix); y++ {
		// numBuf is used to store the number we are currently parsing
		numBuf := ""
		// x is the position in the line
		x := 0
		// gears is a list of all the gears that are touching the current part
		gears := make([]pos, 0)
		// Loop until somthing breaks
		for {
			// Check if we are at the end of the line
			eol := x == len(matrix[y])
			// If we are at the end of a part (end of line or not a number)
			if eol || !isNum(matrix[y][x]) {
				// Check we were on a part before
				if numBuf != "" {
					// Parse the part
					intNum, err := strconv.Atoi(numBuf)
					if err != nil {
						panic(err)
					}
					// Get the parts starting position
					startPos := pos{x - len(numBuf), y}
					// If there were gears touching the part, store it
					if len(gears) > 0 {
						allParts[startPos] = intNum
					}
					// For every gear this part touched, store this part number for that gear
					for _, gear := range unique(gears) {
						if _, ok := allGears[gear]; !ok {
							allGears[gear] = make(map[pos]int)
						}
						allGears[gear][startPos] = intNum
					}

					// Reset buffer state
					numBuf = ""
					gears = make([]pos, 0)
				}
				// Break if we hit end of line
				if eol {
					break
				}
			} else {
				// Add a number to the buffer
				numBuf += string(matrix[y][x])
				// Check for gears in surrounding area
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						if dy == 0 && dx == 0 {
							continue
						}
						if isPossibleGear(matrix, y+dy, x+dx) {
							gears = append(gears, pos{x + dx, y + dy})
						}
					}
				}
			}
			x++
		}
	}
	// Print out total of all parts touching gears
	total := 0
	for _, num := range allParts {
		total += num
	}
	fmt.Println("Total of all parts touching gears:", total)

	// Print out total 'Gear Ratio'
	totalGR := 0
	for _, gear := range allGears {
		if len(gear) > 1 {
			gr := 1
			for _, num := range gear {
				gr *= num
			}
			totalGR += gr
		}
	}
	fmt.Println("Total 'Gear Ratio':", totalGR)
}
