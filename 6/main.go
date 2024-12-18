package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Coordinates of an item
type Position struct {
	Row int
	Col int
}

type Direction struct {
	Vertical   int
	Horizontal int
}

func main() {
	// Directions
	// TODO this could be possibly used to traverse the map in a refactored way
	Up := Direction{-1, 0}
	Down := Direction{1, 0}
	Left := Direction{0, -1}
	Right := Direction{0, 1}

	var floor_map [][]string // 2D array of the map
	var counter int          // Total count of distinct visited places

	var current_guard_position Position // Current position of the guard
	current_guard_direction := Up       // Current guard direction

	// Open the file
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Scan the whole file into 2D array and also find the starting position of the guard
	row := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()           // Get line of input
		words := strings.Split(line, "") // Split line inbetween
		var floor_map_line []string
		col := 0
		for _, v := range words {
			floor_map_line = append(floor_map_line, v)
			if v == "^" {
				current_guard_position = Position{row, col}
			}
			col++
		}
		floor_map = append(floor_map, floor_map_line)
		row++
	}

	for {
		if current_guard_direction == Up {
			// The guard is going up, in for cycle move him up until he hits a #
			for i := current_guard_position.Row; i >= 0; i-- {
				// Update the current guard position, only if the position is not #, if it is #, leave the position but change directions and break
				if floor_map[i][current_guard_position.Col] == "#" {
					current_guard_direction = Right
					break
				} else {
					// Check if you are on the edge of the map, if yes, print result and terminate
					if i-1 < 0 {
						counter++
						fmt.Print(counter)
						os.Exit(0)
					}
					// If the position is not #, and is not X, change it to "X" and update the guards position, if change to X occured, add +1 to the total counter
					if floor_map[i][current_guard_position.Col] != "X" {
						floor_map[i][current_guard_position.Col] = "X"
						counter++
					}
					current_guard_position = Position{i, current_guard_position.Col}
				}
			}
		}
		if current_guard_direction == Down {
			// The guard is going down, in for cycle move him down until he hits a #
			for i := current_guard_position.Row; i < len(floor_map); i++ {
				if floor_map[i][current_guard_position.Col] == "#" {
					current_guard_direction = Left
					break
				} else {
					if i+1 == len(floor_map) {
						counter++
						fmt.Print(counter)
						os.Exit(0)
					}
					if floor_map[i][current_guard_position.Col] != "X" {
						floor_map[i][current_guard_position.Col] = "X"
						counter++
					}
					current_guard_position = Position{i, current_guard_position.Col}
				}
			}
		}
		if current_guard_direction == Left {
			// The guard is going left, in for cycle move him left until he hits a #
			for j := current_guard_position.Col; j >= 0; j-- {
				if floor_map[current_guard_position.Row][j] == "#" {
					current_guard_direction = Up
					break
				} else {
					if j-1 < 0 {
						counter++
						fmt.Print(counter)
						os.Exit(0)
					}
					if floor_map[current_guard_position.Row][j] != "X" {
						floor_map[current_guard_position.Row][j] = "X"
						counter++
					}
					current_guard_position = Position{current_guard_position.Row, j}
				}
			}
		}
		if current_guard_direction == Right {
			// The guard is going right, in for cycle move him left until he hits a #
			for j := current_guard_position.Col; j < len(floor_map[0]); j++ {
				if floor_map[current_guard_position.Row][j] == "#" {
					current_guard_direction = Down
					break
				} else {
					if j+1 == len(floor_map[0]) {
						counter++
						fmt.Print(counter)
						os.Exit(0)
					}
					if floor_map[current_guard_position.Row][j] != "X" {
						floor_map[current_guard_position.Row][j] = "X"
						counter++
					}
					current_guard_position = Position{current_guard_position.Row, j}
				}
			}
		}
	}
}
