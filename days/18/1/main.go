package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Vector struct {
	x int
	y int
}

type Visited map[Vector]int

type Item struct {
	p Vector
	s int
}

type Queue []Item

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseInput(input *[]string) []Vector {

	re := regexp.MustCompile(`^(\d+),(\d+)$`)
	bytes := []Vector{}

	for _, line := range *input {
		matches := re.FindStringSubmatch(line)
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		row := Vector{x, y}
		bytes = append(bytes, row)
	}
	return bytes
}

func buildGrid(size int, bytes *[]Vector, n int) [][]byte {

	grid := [][]byte{}
	usedN := 0
	usedBytes := (*bytes)[:n]

	for y := 0; y <= size; y++ {
		grid = append(grid, []byte{})
		for x := 0; x <= size; x++ {
			foundInBytes := false
			if usedN < n {
				for _, b := range usedBytes {
					if b.x == x && b.y == y {
						foundInBytes = true
						continue
					}
				}
			}

			if foundInBytes {
				grid[y] = append(grid[y], '#')
				usedN++
			} else {
				grid[y] = append(grid[y], '.')
			}

		}
	}
	return grid
}

func draw(grid [][]byte) {
	for _, y := range grid {
		for _, x := range y {
			fmt.Print(string(x))
		}
		fmt.Println()
	}
}

func sort(items []Item) []Item {

	if len(items) <= 1 {
		return items
	}
	pivot := items[0]

	less := []Item{}
	greater := []Item{}
	eq := []Item{}

	for _, item := range items {
		if item.s < pivot.s {
			less = append(less, item)
		} else if item.s > pivot.s {
			greater = append(greater, item)
		} else {
			eq = append(eq, item)
		}
	}

	return append(append(sort(less), eq...), sort(greater)...)
}

func (q *Queue) sort() {
	*q = sort(*q)
}

func (q *Queue) push(item Item) {
	*q = append(*q, item)
	(*q).sort()
}

func (q *Queue) pop() Item {
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func isValidPosition(grid *[][]byte, p Vector) bool {
	return p.x >= 0 && p.x < len(*grid) && p.y >= 0 && p.y < len((*grid)[0]) && (*grid)[p.y][p.x] == '.'
}

func addVector(a Vector, b Vector) Vector {
	return Vector{a.x + b.x, a.y + b.y}
}

func solve(grid *[][]byte, end Vector) (int, []Vector) {

	queue := Queue{Item{Vector{0, 0}, 0}}
	visited := Visited{}

	for len(queue) > 0 {

		current := queue.pop()

		if current.p == end {
			// fmt.Println(current)
			return current.s, []Vector{}
		}

		if visited[current.p] > 0 {
			continue
		}

		for _, v := range []Vector{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			if isValidPosition(grid, addVector(current.p, v)) {
				queue.push(Item{addVector(current.p, v), current.s + 1})
				visited[current.p] = current.s
			}
		}

		// fmt.Println(current)
	}

	return 0, []Vector{}
}

func main() {
	input := readInput()
	bytes := parseInput(&input)
	size := 70
	grid := buildGrid(size, &bytes, 1024)

	fmt.Println(solve(&grid, Vector{size, size}))

	// draw(grid)
}
