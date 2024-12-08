package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var enabled_multi bool = true
	var final_count int = 0
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	str := string(file) // convert content to a 'string'
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don\'t\(\)`)
	matches := re.FindAllString(str, -1)
	re_numbers := regexp.MustCompile(`(\d+),(\d+)`)
	re_multi_enabled := regexp.MustCompile(`do\(\)`)
	re_multi_disabled := regexp.MustCompile(`don\'t\(\)`)
	for _, match := range matches {
		// check if the instruction is do or dont and set the enable_multi and continue
		if re_multi_enabled.FindAllString(match, -1) != nil {
			enabled_multi = true
			continue
		}
		if re_multi_disabled.FindAllString(match, -1) != nil {
			enabled_multi = false
			continue
		}
		// extract numbers if multiplication is enabled
		if enabled_multi {
			numbers_submatch := re_numbers.FindAllString(match, -1)
			if numbers_submatch != nil {
				for _, m := range numbers_submatch {
					stringSlice := strings.Split(m, ",")
					m1, _ := strconv.Atoi(string(stringSlice[0]))
					m2, _ := strconv.Atoi(string(stringSlice[1]))
					final_count += m1 * m2
				}
			}
		}
	}
	fmt.Print(final_count)
}
