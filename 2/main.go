package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Part 1
	var levels [][]string  // Slice of the levels loaded from the text file
	var levels_int [][]int // Slice of the levels loaded from the text file
	var safe_reports int   // Count of safe reports in the input

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Load the lines into slice of slice of strings
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		levels_string := strings.Fields(line)  // Split file inbetween
		levels = append(levels, levels_string) // Get line of input
	}

	// Convert the slice of slice of strings to ints
	for i := range levels {
		var temp_int_levels_slice []int
		for _, v := range levels[i] {
			temp_level, _ := strconv.Atoi(v)
			temp_int_levels_slice = append(temp_int_levels_slice, temp_level)
		}
		levels_int = append(levels_int, temp_int_levels_slice)
	}

	// Find all safe reports - the ints (levels) in a report must be either ascending
	// or descending and each next level must be bigger/smaller by 1-3 than the prior
	for i := range levels_int {
		// Temporary count of OK comparisons in each report
		temp_count := 0
		// Is the range ascending or descending?
		// If both of these are true, the report is unsafe as the levels cannot be ascending and descending at once
		ascending := false
		descending := false
		for j, v := range levels_int[i] {
			if j == len(levels_int[i])-1 {
				break
			}
			next_value := levels_int[i][j+1]
			if v < next_value {
				ascending = true
			}
			if v > next_value {
				descending = true
			}
			if v == next_value {
				ascending = true
				descending = true
			}
			// Check that the next level is ascending/descending and at most by 1-3
			if (math.Abs(float64(v)-float64(next_value)) <= 3 && math.Abs(float64(v)-float64(next_value)) >= 1) && ascending != descending {
				temp_count++
			}
		}
		// This means that all levels in a report are OK
		if temp_count == len(levels_int[i])-1 {
			safe_reports++
		}
	}
	// Part 1 result
	fmt.Print(safe_reports)
}
