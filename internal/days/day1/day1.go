package day1

import (
	"bufio"
	"fmt"
	"strings"
)

// Day1 implements the Day interface
type Day1 struct{}

func (d *Day1) Run(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	sum := 0
	for scanner.Scan() {
		// Example logic: Sum all integers in the input
		var num int
		fmt.Sscanf(scanner.Text(), "%d", &num)
		sum += num
	}
	return fmt.Sprintf("Day 1 result: %d", sum)
}
