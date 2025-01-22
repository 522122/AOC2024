package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	x int
	y int
}

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

var numericKeyboard = [][]string{
	{"7", "8", "9"},
	{"4", "5", "6"},
	{"1", "2", "3"},
	{"", "0", "A"},
}

var directionalKeyboard = [][]string{
	{"", "^", "A"},
	{"<", "v", ">"},
}

func possibleCommands(desired string, current string, kayboardMap [][]string) [][]string {
	keys := [][]string{}
	var position Vector

	var desiredPosition Vector
	for y, row := range kayboardMap {
		for x, key := range row {
			if key == desired {
				desiredPosition = Vector{x, y}
			}
			if key == current {
				position = Vector{x, y}
			}
		}
	}

	if kayboardMap[position.y][position.x] == desired {
		return keys
	}

	var recursion func(from Vector, to Vector, innerKeys []string)
	recursion = func(from Vector, to Vector, innerKeys []string) {

		if kayboardMap[from.y][from.x] == desired {
			keys = append(keys, innerKeys)
			return
		}

		if from.x > to.x && kayboardMap[from.y][from.x-1] != "" {
			newKeys := append(innerKeys, "<")
			recursion(Vector{from.x - 1, from.y}, to, newKeys)
		}
		if from.y < to.y && kayboardMap[from.y+1][from.x] != "" {
			newKeys := append(innerKeys, "v")
			recursion(Vector{from.x, from.y + 1}, to, newKeys)
		}
		if from.y > to.y && kayboardMap[from.y-1][from.x] != "" {
			newKeys := append(innerKeys, "^")
			recursion(Vector{from.x, from.y - 1}, to, newKeys)
		}
		if from.x < to.x && kayboardMap[from.y][from.x+1] != "" {
			newKeys := append(innerKeys, ">")
			recursion(Vector{from.x + 1, from.y}, to, newKeys)
		}
	}

	recursion(position, desiredPosition, []string{})

	return keys
}

func commandsToNumeric(desired string, current string) []string {
	out := []string{}
	commands := possibleCommands(desired, current, numericKeyboard)
	if len(commands) == 0 {
		return append(out, "")
	} else {
		for _, c := range commands {
			out = append(out, strings.Join(c, ""))
		}
		return out
	}
}

func createNumPaths() map[string][]string {
	out := map[string][]string{}
	for _, topFrom := range []string{"A", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"} {
		for _, topTo := range []string{"A", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"} {
			paths := commandsToNumeric(topTo, topFrom)
			key := topFrom + "-" + topTo
			out[key] = append(out[key], paths...)
		}
	}

	return out
}

func countSteps(start map[string]int, rules map[string]string, n int) int {
	for x := 0; x < n; x++ {
		change := map[string]int{}
		for fromTo, count := range start {
			if count == 0 {
				continue
			}
			next := "A" + rules[fromTo] + "A"

			for i := 0; i < len(next)-1; i++ {
				key := string(next[i]) + "-" + string(next[i+1])
				change[key] += count
			}
		}
		start = change
	}

	sum := 0
	for _, count := range start {
		sum += count
	}
	return sum
}

func bestNumericPath(start string, rules map[string]string, numPaths map[string][]string) string {
	start = "A" + start
	out := "A"
	for i := 0; i < len(start)-1; i++ {
		key := string(start[i]) + "-" + string(start[i+1])

		best := ""
		bestSteps := 0
		for _, p := range numPaths[key] {
			testMap := pathToMap(p)
			steps := countSteps(testMap, rules, 2)
			if steps < bestSteps || bestSteps == 0 {
				bestSteps = steps
				best = p
			}
		}

		out += best + string("A")
	}
	return out
}

func pathToMap(path string) map[string]int {
	out := map[string]int{}
	for i := 0; i < len(path)-1; i++ {
		out[string(path[i])+"-"+string(path[i+1])]++
	}
	return out
}

func main() {
	input := readInput()

	rules := map[string]string{
		"A-^": "<",
		"A-v": "<v",
		"A->": "v",
		"A-<": "v<<",
		"A-A": "",

		"v-A": "^>",
		"v-v": "",
		"v->": ">",
		"v-<": "<",
		"v-^": "^",

		">-A": "^",
		">-v": "<",
		">->": "",
		">-<": "<<",
		">-^": "<^",

		"<-A": ">>^",
		"<-v": ">",
		"<->": ">>",
		"<-<": "",
		"<-^": ">^",

		"^->": "v>",
		"^-<": "v<",
		"^-v": "v",
		"^-A": ">",
		"^-^": "",
	}

	numPaths := createNumPaths()
	sum := 0
	for _, line := range input {
		bestPath := bestNumericPath(line, rules, numPaths)
		steps := countSteps(pathToMap(bestPath), rules, 25)

		intLine, _ := strconv.Atoi(line[0 : len(line)-1])
		fmt.Println(intLine, steps)
		sum += intLine * steps

	}

	fmt.Println(sum)
}
