package main

import (
	"bufio"
	"fmt"
	"os"
)

// Coordinates of "X"
type XPosition struct {
	X int
	Y int
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

	var map_of_xs []XPosition
	// Create a map of positions of all Xs
	for i, row := range character_matrix {
		for j, character := range row {
			if character == "X" {
				map_of_xs = append(map_of_xs, XPosition{i, j})
			}
		}
	}

	// Go through the rows on Xpositions in all directions
	for _, position := range map_of_xs {
		// Horizontal direction
		total_xmas_counter += sum_up_horizontal_positive_direction(position, character_matrix)
		total_xmas_counter += sum_up_horizontal_negative_direction(position, character_matrix)

	}
	fmt.Print(total_xmas_counter)
}

func determine_found_characters(position string, m_truth_val *bool, a_truth_val *bool, s_truth_val *bool) bool {
	// Determine if the character found is one of M A S
	if position == "M" {
		*m_truth_val = true
	}
	if position == "A" && *m_truth_val {
		*a_truth_val = true
	}
	if position == "S" && *m_truth_val && *a_truth_val {
		*s_truth_val = true
	}
	if *m_truth_val && *a_truth_val && *s_truth_val {
		// We found Xmas - return 1 to be summed up to the total
		return true
	}
	return false
}

func sum_up_horizontal_positive_direction(position XPosition, matrix [][]string) int {
	found_m, found_a, found_s := false, false, false
	for row_y := position.Y + 1; row_y < len(matrix); row_y++ {
		if determine_found_characters(matrix[position.X][row_y], &found_m, &found_a, &found_s) {
			return 1
		}
	}
	return 0
}

func sum_up_horizontal_negative_direction(position XPosition, matrix [][]string) int {
	found_m, found_a, found_s := false, false, false
	for row_y := position.Y - 1; row_y > 0; row_y-- {
		if determine_found_characters(matrix[position.X][row_y], &found_m, &found_a, &found_s) {
			return 1
		}
	}
	return 0
}

func sum_up_vertical_positive_direction(position XPosition, matrix [][]string) int {
	found_m, found_a, found_s := false, false, false
	for row_y := position.Y - 1; row_y > 0; row_y-- {
		if determine_found_characters(matrix[position.X][row_y], &found_m, &found_a, &found_s) {
			return 1
		}
	}
	return 0
}
