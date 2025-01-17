package utils

import (
	"bufio"
	"os"
)

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func sort[T any](items []T, comparator func(t1 T, t2 T) int) []T {

	if len(items) <= 1 {
		return items
	}

	pivot := items[0]

	var greater, less, equal []T

	for _, item := range items {

		result := comparator(item, pivot)
		if result > 0 {
			greater = append(greater, item)
		} else if result < 0 {
			less = append(less, item)
		} else {
			equal = append(equal, item)
		}
	}

	return append(append(sort(less, comparator), equal...), sort(greater, comparator)...)
}

func filter[T any](items []T, predicate func(t T) bool) []T {
	var out []T
	for _, item := range items {
		if predicate(item) {
			out = append(out, item)
		}
	}
	return out
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
