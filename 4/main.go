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

	// Create a 2D array to hold the result
	var result [][]string

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // Read the current line
		var chars []string
		for _, char := range line {
			chars = append(chars, string(char)) // Convert each rune to a string
		}
		result = append(result, chars)
	}

	// Check for errors in reading the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var map_of_xs []XPosition
	// Create a map of positions of all Xs
	for i, row := range result {
		for j, character := range row {
			if character == "X" {
				map_of_xs = append(map_of_xs, XPosition{i, j})
			}
		}
	}

	fmt.Print(map_of_xs)

	// Go through the rows in horizontal positive direction
	for _, row := range result {
		for j, character := range row {
			if j+3 == len(row) {
				// No more to be found in this row on both sides
				break
			}
			if character == "X" {
				// We found X, now check all directions
				found_m := false
				found_a := false
				found_s := false
				// Check that we find M,A,S in this succession in horizontal positive direction
				for row_i := j; row_i < len(row); row_i++ {
					if row[row_i] == "M" {
						found_m = true
					}
					if row[row_i] == "A" && found_m {
						found_a = true
					}
					if row[row_i] == "S" && found_m && found_a {
						found_s = true
					}
					if found_m && found_a && found_s {
						// We found Xmas
						total_xmas_counter++
						break
					}

				}
			}
		}

	}
	fmt.Print(total_xmas_counter)
}
