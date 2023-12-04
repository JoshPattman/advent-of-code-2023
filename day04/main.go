package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// Find the intersection of two slices
func intersection[T comparable](a, b []T) []T {
	out := make([]T, 0)
	for _, x := range a {
		for _, y := range b {
			if x == y {
				out = append(out, x)
			}
		}
	}
	return out
}

// Check a card, and return how many cards we end up checking
func checkCard(numIntersections map[int]int, id int) int {
	cardWins := numIntersections[id]
	totalChecked := 1
	for i := 0; i < cardWins; i++ {
		totalChecked += checkCard(numIntersections, id+i+1)
	}
	return totalChecked
}

// Main entrypoint
func main() {
	// Split input into lines, where each line is a card
	cards := strings.Split(input, "\n")
	// Remeber the total 'points' for each card
	total := 0
	// Remeber the number of intersections (wins) that each card has
	numIntersections := make(map[int]int)
	// Loop over each card
	for _, card := range cards {
		// Split the card into its data and id
		idData := strings.Split(card, ": ")
		idParts := strings.Split(idData[0], " ")
		id, _ := strconv.Atoi(idParts[len(idParts)-1])
		data := idData[1]
		// Split the data into the winning numbers and my numbers
		cardParts := strings.Split(data, " | ")
		winNumsStr, myNumsStr := cardParts[0], cardParts[1]
		winNums := make([]int, 0)
		myNums := make([]int, 0)
		for _, numStr := range strings.Split(winNumsStr, " ") {
			if numStr == "" {
				continue
			}
			num, _ := strconv.Atoi(numStr)
			winNums = append(winNums, num)
		}
		for _, numStr := range strings.Split(myNumsStr, " ") {
			if numStr == "" {
				continue
			}
			num, _ := strconv.Atoi(numStr)
			myNums = append(myNums, num)
		}
		// Find the intersection of the winning numbers and my numbers
		intersectNums := intersection(winNums, myNums)
		// Store the number of intersections
		numIntersections[id] = len(intersectNums)
		// Increment the total points
		total += int(math.Pow(2, float64(len(intersectNums)-1)))
	}

	// Check (recursive) each card to find out how many cards in total we checked
	numCardsChecked := 0
	for ci := range cards {
		numCardsChecked += checkCard(numIntersections, ci)
	}

	fmt.Println("Total points:", total)
	fmt.Println("Total cards checked:", numCardsChecked)
}
