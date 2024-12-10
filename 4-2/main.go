package main

import (
	"bufio"
	"fmt"
	"os"
)

// Coordinates of "A"
type Position struct {
	Row int
	Col int
}

func main() {
	var total_xmas_counter int = 0
	// Open the file
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a 2D array to hold the character_matrix
	var character_matrix [][]string

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // Read the current line
		var chars []string
		for _, char := range line {
			chars = append(chars, string(char)) // Convert each rune to a string
		}
		character_matrix = append(character_matrix, chars)
	}

	// Check for errors in reading the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var map_of_as []Position
	// Create a map of positions of all Xs
	for i, row := range character_matrix {
		for j, character := range row {
			if character == "A" {
				map_of_as = append(map_of_as, Position{i, j})
			}
		}
	}

	// Pick all A as starting position and check the X-mas pattern
	for _, position := range map_of_as {
		r := position.Row
		c := position.Col
		if r-1 >= 0 && r+1 < len(character_matrix) && c-1 >= 0 && c+1 < len(character_matrix[0]) {
			var diag1 string
			var diag2 string

			diag1 = character_matrix[r-1][c-1] + character_matrix[r][c] + character_matrix[r+1][c+1]
			diag2 = character_matrix[r-1][c+1] + character_matrix[r][c] + character_matrix[r+1][c-1]

			if (diag1 == "MAS" || diag1 == "SAM") && (diag2 == "MAS" || diag2 == "SAM") {
				total_xmas_counter++
			}
		}
	}
	fmt.Print(total_xmas_counter)
}
