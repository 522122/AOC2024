package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Maze [][]byte
type Visited map[Vector]bool

type Vector struct {
	x int
	y int
}

type Item struct {
	position  Vector
	direction Vector
	cost      int
}

type Queue []*Item

func (q *Queue) Push(x any) {
	*q = append(*q, x.(*Item))
}

func (q *Queue) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *Queue) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	*q = old[0 : n-1]
	return item
}

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) Less(i, j int) bool {
	return (*q)[i].cost < (*q)[j].cost
}

var directions = []Vector{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

const (
	WALL  = '#'
	SPACE = '.'
	START = 'S'
	END   = 'E'
)

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseInput(lines *[]string) Maze {
	out := Maze{}
	for _, line := range *lines {
		out = append(out, []byte(line))
	}
	return out
}

func isInPath(path []Vector, v Vector) bool {
	for _, p := range path {
		if p.x == v.x && p.y == v.y {
			return true
		}
	}
	return false
}

func draw(maze *Maze, path []Vector) {
	custoMpathSum := 0
	for y, line := range *maze {
		for x, cell := range line {
			if isInPath(path, Vector{x, y}) {
				fmt.Print("O")
				custoMpathSum++
			} else {
				fmt.Print(string(cell))
			}
		}
		fmt.Print("\n")
	}
	fmt.Println(custoMpathSum)
}

func findInMaze(maze *Maze, object byte) Vector {
	for y, line := range *maze {
		for x, cell := range line {
			if cell == object {
				return Vector{x, y}
			}
		}
	}
	return Vector{-1, -1}
}

func addVectors(a Vector, b Vector) Vector {
	return Vector{a.x + b.x, a.y + b.y}
}

func dot(a Vector, b Vector) int {
	return a.x*b.x + a.y*b.y
}

func solve(maze *Maze) ([]Vector, int) {
	queue := &Queue{}
	heap.Init(queue)

	var start Vector = findInMaze(maze, START)
	var end Vector = findInMaze(maze, END)
	visited := Visited{}
	parent := map[Vector]Vector{}

	heap.Push(queue, &Item{start, Vector{1, 0}, 0})

	for queue.Len() > 0 {
		item := heap.Pop(queue).(*Item)

		if item.position == end {
			// reconstruct path
			path := []Vector{}

			for p := end; p != start; p = parent[p] {
				path = append([]Vector{p}, path...)
			}
			path = append([]Vector{start}, path...)
			return path, item.cost
		}

		for _, direction := range directions {
			newPosition := addVectors(item.position, direction)

			if (*maze)[newPosition.y][newPosition.x] == WALL {
				continue
			}

			if _, ok := visited[newPosition]; ok {
				continue
			}

			switch dot(direction, item.direction) {
			case 0:
				heap.Push(queue, &Item{item.position, direction, item.cost + 1000})
			case 1:
				heap.Push(queue, &Item{newPosition, direction, item.cost + 1})
				visited[newPosition] = true
			case -1:
				continue
			}

			parent[newPosition] = item.position
		}
	}

	return nil, -1
}

func main() {
	lines := readInput()

	var maze Maze = parseInput(&lines)

	path, score := solve(&maze)

	fmt.Printf("Path: %v\n", path)

	draw(&maze, path)

	fmt.Println(score)
}
