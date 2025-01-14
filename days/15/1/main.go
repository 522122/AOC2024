package main

import (
	"bufio"
	"fmt"
	"os"
)

type Vector struct {
	x int
	y int
}

type Map [][]byte

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseInput(input []string) (Map, []string) {
	splitIndex := -1
	for i, line := range input {
		if line == "" {
			splitIndex = i
		}
	}

	tmp := input[:splitIndex]
	warehouse := make(Map, len(tmp))
	for i, line := range tmp {
		warehouse[i] = []byte(line)
	}

	return warehouse, input[splitIndex+1:]
}

func findStart(warehouse *Map) Vector {
	for y, line := range *warehouse {
		for x, char := range line {
			if char == '@' {
				return Vector{x, y}
			}
		}
	}
	return Vector{-1, -1}
}

func toVector(instruction string) Vector {
	switch instruction {
	case "^":
		return Vector{0, -1}
	case "v":
		return Vector{0, 1}
	case "<":
		return Vector{-1, 0}
	case ">":
		return Vector{1, 0}
	}
	return Vector{0, 0}
}

func moveObject(warehouse *Map, from *Vector, to Vector) {
	tmpTo := (*warehouse)[to.y][to.x]
	tmpFrom := (*warehouse)[from.y][from.x]

	(*warehouse)[to.y][to.x] = tmpFrom
	(*warehouse)[from.y][from.x] = tmpTo
	from.x = to.x
	from.y = to.y
}

func canPush(warehouse *Map, position Vector, direction *Vector) (bool, int) {
	n := 0
	for {
		object := (*warehouse)[position.y][position.x]
		switch object {
		case '.':
			return true, n
		case '#':
			return false, 0
		case 'O':
			position.x += direction.x
			position.y += direction.y
			n++
		}
	}
}

func playMove(warehouse *Map, instruction string, position *Vector) {

	direction := toVector(instruction)

	destination := Vector{position.x + direction.x, position.y + direction.y}

	object := (*warehouse)[destination.y][destination.x]

	switch object {
	case '.':
		moveObject(warehouse, position, destination)
	case '#':
	case 'O':
		nextPosition := Vector{position.x + direction.x * 2, position.y + direction.y * 2}
		nextObject := (*warehouse)[nextPosition.y][nextPosition.x]

		fmt.Println("next object: ", string(nextObject), " at ", nextPosition)

		switch nextObject {
		case '#':
			break
		case 'O':
			canPush, n := canPush(warehouse, destination, &direction)
			if !canPush {
				break
			}
			fmt.Println("Can push ", n, " objects")
			for i := n ; i > 0; i-- {
				moveFrom := Vector{destination.x + direction.x * (i - 1), destination.y + direction.y * (i - 1)}
				moveTo := Vector{destination.x + direction.x * i, destination.y + direction.y * i}
				fmt.Println("Move ", string((*warehouse)[moveFrom.y][moveFrom.x]) ," from ", moveFrom, " to ", moveTo)
				moveObject(warehouse, &moveFrom, moveTo)
			}
			moveObject(warehouse, position, destination)
		case '.':
			moveObject(warehouse, &Vector{destination.x, destination.y}, nextPosition)
			moveObject(warehouse, position, destination)
			break
		}
	}
}

func gpsCoords(position Vector) int {
	return 100 * position.y + position.x
}

func sumOfGpsCoords(warehouse *Map) int {
	sum := 0
	for y, line := range *warehouse {
		for x, char := range line {
			if char == 'O' {
				sum += gpsCoords(Vector{x, y})
			}
		}
	}

	return sum
}

func main() {

	input := readInput()

	warehouse, instructions := parseInput(input)
	start := findStart(&warehouse)

	for _, instructionLine := range instructions {
		for _, instruction := range instructionLine {
			playMove(&warehouse, string(instruction), &start)

			fmt.Println("Move ", string(instruction), ":")
			for _, line := range warehouse {
				fmt.Println(string(line))
			}
			fmt.Println()
		}
	}

	fmt.Println(sumOfGpsCoords(&warehouse))

}