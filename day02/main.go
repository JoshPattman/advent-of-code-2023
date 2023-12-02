package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// The set of constraints needed for the puzzle
var constraintSet = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// A game represents a game of cubes
type Game struct {
	ID                 int            // The id of the game
	MinCubeConstraints map[string]int // A map of constraints for the minimum number of cubes of each color
}

// Parse a game from a line. This will calculate relevant constraints and save them to the game
func ParseGame(line string) (*Game, error) {
	// Read the game ID
	idData := strings.Split(line, ": ")
	id, err := strconv.Atoi(strings.Split(idData[0], " ")[1])
	if err != nil {
		return nil, err
	}
	// Create a list to store the minimum constraints for each color
	minCubeConstraints := make(map[string]int)
	// Loop over each turn in the game
	turns := strings.Split(idData[1], "; ")
	for _, turn := range turns {
		// Loop over every color that appeared in this turn
		turnColorsCounts := strings.Split(turn, ", ")
		for _, turnColorCount := range turnColorsCounts {
			// Get the color and count
			turnColorCountParts := strings.Split(turnColorCount, " ")
			turnCount, err := strconv.Atoi(turnColorCountParts[0])
			if err != nil {
				return nil, err
			}
			turnColor := turnColorCountParts[1]
			// Check if this a a more constrictive constraint than the current one, if so, replace it
			if currentMin, ok := minCubeConstraints[turnColor]; (ok && turnCount > currentMin) || !ok {
				minCubeConstraints[turnColor] = turnCount
			}
		}
	}
	// Return the game
	return &Game{
		ID:                 id,
		MinCubeConstraints: minCubeConstraints,
	}, nil
}

// Check if a game is possible given a set of colored cubes that are in play
func (g *Game) CheckPossible(knownCubeCounts map[string]int) bool {
	for conCol, conCount := range g.MinCubeConstraints {
		if knownCubeCounts[conCol] < conCount {
			return false
		}
	}
	return true
}

// Calculate the 'power' of a game: the product of the minimum number of cubes of each color
func (g *Game) Power() int {
	v := 1
	for _, conCount := range g.MinCubeConstraints {
		v *= conCount
	}
	return v
}

// Main function
func main() {
	// Split the input file into lines
	lines := strings.Split(input, "\n")
	// Parse each line into a game
	games := make([]*Game, 0)
	for _, line := range lines {
		if game, err := ParseGame(line); err != nil {
			panic(err)
		} else {
			games = append(games, game)
		}
	}

	// Calculate the total power, and total id of possible games
	totalId := 0
	totalPower := 0
	for _, game := range games {
		power := game.Power()
		totalPower += power

		possible := game.CheckPossible(constraintSet)
		if possible {
			totalId += game.ID
		}
		fmt.Printf("Game %3d: [Power: %4d] [Possible: %5t]\n", game.ID, power, possible)
	}

	// Print the results
	fmt.Println()
	fmt.Printf("Total ID of possible games: %d\n", totalId)
	fmt.Printf("Total power of all games: %d\n", totalPower)
}
