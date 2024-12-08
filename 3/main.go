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
	var final_count int = 0
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	str := string(file) // convert content to a 'string'
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllString(str, -1)
	re = regexp.MustCompile(`(\d+),(\d+)`)
	for _, match := range matches {
		// extract numbers
		submatch := re.FindAllString(match, -1)
		for _, m := range submatch {
			stringSlice := strings.Split(m, ",")
			m1, _ := strconv.Atoi(string(stringSlice[0]))
			m2, _ := strconv.Atoi(string(stringSlice[1]))
			final_count += m1 * m2
		}
	}
	fmt.Print(final_count)
}
