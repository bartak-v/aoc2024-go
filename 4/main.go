package main

import (
	"bufio"
	"fmt"
	"os"
)

// Coordinates of "X"
type XPosition struct {
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
		total_xmas_counter += sum_horizontal_positive_direction(position, character_matrix)
		total_xmas_counter += sum_horizontal_negative_direction(position, character_matrix)
		total_xmas_counter += sum_vertical_negative_direction(position, character_matrix)
		total_xmas_counter += sum_vertical_positive_direction(position, character_matrix)
		total_xmas_counter += sum_right_diagonal_positive_direction(position, character_matrix)
		total_xmas_counter += sum_right_diagonal_negative_direction(position, character_matrix)
		total_xmas_counter += sum_left_diagonal_positive_direction(position, character_matrix)
		total_xmas_counter += sum_left_diagonal_negative_direction(position, character_matrix)
	}
	fmt.Print(total_xmas_counter)
}

func determine_found_characters(position string, m_truth_val *bool, a_truth_val *bool, s_truth_val *bool) bool {
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
		return true
	}
	return false
}

func sum_horizontal_positive_direction(position XPosition, matrix [][]string) int {
	found_m, found_a, found_s := false, false, false
	for col := position.Col; col < len(matrix[0]); col++ {
		if determine_found_characters(matrix[position.Row][col], &found_m, &found_a, &found_s) {
			return 1
		}
	}
	return 0
}

func sum_horizontal_negative_direction(position XPosition, matrix [][]string) int {
	found_m, found_a, found_s := false, false, false
	for col := position.Col; col > 0; col-- {
		if determine_found_characters(matrix[position.Row][col], &found_m, &found_a, &found_s) {
			return 1
		}
	}
	return 0
}

func sum_vertical_positive_direction(position XPosition, matrix [][]string) int {
	found_m, found_a, found_s := false, false, false
	for row := position.Row; row < len(matrix); row++ {
		if determine_found_characters(matrix[row][position.Col], &found_m, &found_a, &found_s) {
			return 1
		}
	}
	return 0
}

func sum_vertical_negative_direction(position XPosition, matrix [][]string) int {
	found_m, found_a, found_s := false, false, false
	for row := position.Row - 1; row > 0; row-- {
		if determine_found_characters(matrix[row][position.Col], &found_m, &found_a, &found_s) {
			return 1
		}
	}
	return 0
}

func sum_right_diagonal_positive_direction(position XPosition, matrix [][]string) int {
	found_m, found_a, found_s := false, false, false
	for r, c := position.Row, position.Col; r < len(matrix) && c < len(matrix[0]); r, c = r+1, c+1 {
		if determine_found_characters(matrix[r][c], &found_m, &found_a, &found_s) {
			return 1
		}
	}
	return 0
}

func sum_right_diagonal_negative_direction(position XPosition, matrix [][]string) int {
	found_m, found_a, found_s := false, false, false
	for r, c := position.Row, position.Col; r > 0 && c < len(matrix[0]); r, c = r-1, c+1 {
		if determine_found_characters(matrix[r][c], &found_m, &found_a, &found_s) {
			return 1
		}
	}
	return 0
}

func sum_left_diagonal_negative_direction(position XPosition, matrix [][]string) int {
	found_m, found_a, found_s := false, false, false
	for r, c := position.Row, position.Col; r > 0 && c > 0; r, c = r-1, c-1 {
		if determine_found_characters(matrix[r][c], &found_m, &found_a, &found_s) {
			return 1
		}
	}
	return 0
}

func sum_left_diagonal_positive_direction(position XPosition, matrix [][]string) int {
	found_m, found_a, found_s := false, false, false
	for r, c := position.Row, position.Col; r < len(matrix) && c > 0; r, c = r+1, c-1 {
		if determine_found_characters(matrix[r][c], &found_m, &found_a, &found_s) {
			return 1
		}
	}
	return 0
}
