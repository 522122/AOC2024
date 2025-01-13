package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	wall = '#'
	track = '.'
	start = 'S'
	end = 'E'
)

var directions = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

type Map []string

type Cheat struct {
	from [2]int
	to [2]int
	pico int
}

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func findStart(raceMap Map) (int, int) {
	for r, line := range raceMap {
		for c, char := range line {
			if char == start {
				return r, c
			}
		}
	}
	return -1, -1
}

func isInRange(r int, c int, raceMap Map) bool {
	return r >= 0 && r < len(raceMap) && c >= 0 && c < len(raceMap[0])
}

func complete(raceMap Map) (int, int, int, map[string]Cheat, [][2]int) {
	currentR, currentC := findStart(raceMap)
	pico := 0
	visitedMap := map[string]bool{}
	visitedArr := [][2]int{}
	cheatsMap := map[string]Cheat{}

	for {
		key := fmt.Sprintf("%d,%d", currentR, currentC)
		visitedMap[key] = true
		visitedArr = append(visitedArr, [2]int{currentR, currentC})

		for _, d := range directions {
			cheatR := currentR + d[0] * 2
			cheatC := currentC + d[1] * 2

			key := fmt.Sprintf("%d,%d,%d,%d", currentR, currentC, cheatR, cheatC)

			visitedCompareKey := fmt.Sprintf("%d,%d", cheatR, cheatC)

			if (isInRange(cheatR, cheatC, raceMap) && !visitedMap[visitedCompareKey]) {
				canCheat := (raceMap[cheatR][cheatC] == track || raceMap[cheatR][cheatC] == end)
				if canCheat {
					cheatsMap[key] = Cheat{
						from: [2]int{currentR, currentC},
						to: [2]int{cheatR, cheatC},
						pico: pico,
					}
				}
			}
		}

		for i, d := range directions {
			nR := currentR + d[0]
			nC := currentC + d[1]
			key := fmt.Sprintf("%d,%d", nR, nC)

			if (raceMap[nR][nC] == track || raceMap[nR][nC] == end) && !visitedMap[key] {
				currentR = nR
				currentC = nC
				pico++
				break
			}

			if i == len(directions) - 1 {
				fmt.Println("Stuck")
				return -1, -1, pico, cheatsMap, visitedArr
			}
		}
		if raceMap[currentR][currentC] == end {
			visitedArr = append(visitedArr, [2]int{currentR, currentC})
			break
		}
	}
	return currentR, currentC, pico, cheatsMap, visitedArr
}

func skippedWithCheat(cheat Cheat, visited [][2]int) (int) {
	cheatStartIndex := -1
	cheatEndIndex := -1

	for i, v := range visited {
		if v[0] == cheat.from[0] && v[1] == cheat.from[1] {
			cheatStartIndex = i
			break
		}
	}

	for i := cheatStartIndex; i < len(visited); i++ {
		if visited[i][0] == cheat.to[0] && visited[i][1] == cheat.to[1] {
			cheatEndIndex = i
			break
		}
	}
	
	skipped := cheatEndIndex - cheatStartIndex

	return skipped
}

func main() {

	raceMap := readInput()

	r, c, pico, cheats, visited := complete(raceMap)

	fmt.Println(r, c, pico, len(cheats))

	var sumMap = make(map[string]int)

	answerSum := 0

	for _, cheat := range cheats {
		saved := skippedWithCheat(cheat, visited)

		key := fmt.Sprintf("%d", saved - 2)

		if _, ok := sumMap[key]; !ok {
			sumMap[key] = 0
		}
		
		sumMap[key]++

		if saved - 2 >= 100 {
			answerSum++
		}
	}

	fmt.Printf("%d\n",answerSum)
}
