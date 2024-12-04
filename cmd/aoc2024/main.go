package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Ihce/aoc2024/internal/days/day1" // Import day1
)

// Define the interface each day should implement
type Day interface {
	Run(input string) string
}

// A registry to hold the days
var days = map[string]Day{}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: aoc2024 <day>")
		os.Exit(1)
	}

	day := os.Args[1]
	d, ok := days[day]
	if !ok {
		fmt.Printf("Day %s not implemented\n", day)
		os.Exit(1)
	}

	// Load the input file for the day
	inputFile := filepath.Join("internal", "days", day, "input.txt")
	input, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file for day %s: %v\n", day, err)
		os.Exit(1)
	}

	// Run the day's implementation
	result := d.Run(string(input))
	fmt.Println(result)
}

func init() {
	// Register each day here
	days["day1"] = &day1.Day1{}
	// Add registrations for other days as implemented
}
