package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseInput(data []string) ([]string, []string) {
	designs := []string{}

	towels := strings.Split(data[0], ", ")

	for _, line := range data[2:] {
		designs = append(designs, line)
	}

	return towels, designs
}

func memo[I comparable, R any](fn func(i I) R) func(i I) R {
	var cache = map[I]R{}
	return func(match I) R {
		if c, ok := cache[match]; ok {
			return c
		}
		result := fn(match)
		cache[match] = result
		return result
	}
}

func solve(towels *[]string, design string) int {

	var memoized func(match string) int
	memoized = memo(func(match string) int {
		if len(match) == 0 {
			return 1
		}

		count := 0
		for _, towel := range *towels {
			if strings.HasPrefix(match, towel) {
				count += memoized(match[len(towel):])
			}
		}
		// fmt.Println(match, count)
		return count
	})

	return memoized(design)

}

func main() {
	input := readInput()

	var towels []string
	var designs []string
	var start time.Time

	towels, designs = parseInput(input)

	allSum := 0
	solvable := 0

	start = time.Now()

	for _, d := range designs {
		if possibilities := solve(&towels, d); possibilities > 0 {
			solvable++
			allSum += possibilities
		}

	}

	end := time.Now()

	fmt.Println("Solvable:", solvable, "| All ways:", allSum, "| Execution time:", end.Sub(start))

}
