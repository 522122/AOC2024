package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	width = 101 // 11 // 101
	height = 103 // 7 // 103
)

type Vector struct {
	x, y int
}

type Robot struct {
	Position Vector
	Velocity Vector
	Time int
}

type Floor [height][width][]Robot

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func removeRobot(floor *Floor, robot Robot) {
	for i, r := range floor[robot.Position.y][robot.Position.x] {
		if r == robot {
			floor[robot.Position.y][robot.Position.x] = append(floor[robot.Position.y][robot.Position.x][:i], floor[robot.Position.y][robot.Position.x][i+1:]...)
			break
		}
	}
}

func addRobot(floor *Floor, robot Robot) {
	floor[robot.Position.y][robot.Position.x] = append(floor[robot.Position.y][robot.Position.x], robot)
}


func simulate(floor *Floor, time int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			robots := append([]Robot{}, floor[y][x]...)
			for _, robot := range robots {

				// retarded way to avoid double moving, but I can't do any better now
				if robot.Time == time {
					continue
				}

				removeRobot(floor, robot)

				robot.Position.x += robot.Velocity.x
				robot.Position.y += robot.Velocity.y
				robot.Time  = time

				if (robot.Position.x < 0) {
					robot.Position.x = width + robot.Position.x
				}
				if (robot.Position.x >= width) {
					robot.Position.x = robot.Position.x - width
				}
				if (robot.Position.y < 0) {
					robot.Position.y = height + robot.Position.y
				}
				if (robot.Position.y >= height) {
					robot.Position.y = robot.Position.y - height
				}

				addRobot(floor, robot)
			}
		}
	}
}

func draw(floor *Floor, time int, file *os.File) {
	for y, lines := range floor {
		for x := range lines {
			if len(floor[y][x]) > 0 {
				file.WriteString("#")
			} else {
				file.WriteString(".")
			}
		}
		file.WriteString("\n")
	}
	file.WriteString(fmt.Sprintf("%d\n", time))
}

func robotsPerQuadrant(floor *Floor) (int, int, int, int) {
	var q1, q2, q3, q4 int

	middleX := width/2
	middleY := height/2

	for y := 0; y < height; y++ {
		if y == middleY {
			continue
		}
		for x := 0; x < width; x++ {
			if x == middleX {
				continue
			}
			if x < width/2 && y < height/2 {
				q1 += len(floor[y][x])
			} else if x >= width/2 && y < height/2 {
				q2 += len(floor[y][x])
			} else if x < width/2 && y >= height/2 {
				q3 += len(floor[y][x])
			} else {
				q4 += len(floor[y][x])
			}
		}
	}

	return q1, q2, q3, q4
}

func checkTreePattern(floor *Floor) bool {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if len(floor[y][x]) > 0 {

				if x + 8 < width {
					pattern := true
					for i :=1; i<=8;i++ {
						if len(floor[y][x+i]) == 0 {
							pattern = false
							break
						}
					}
					if pattern {
						return true
					}
				}
			}
		}
	}
	return false
}

func main() {

	var input []string = readInput()
	robots := []Robot{}
	floor := Floor{}

	for _, line := range input {

		re := regexp.MustCompile(`[-]?\d+`)
		numbers := re.FindAllString(line, -1)

		p1, _ := strconv.Atoi(numbers[0])
		p2, _ := strconv.Atoi(numbers[1])
		v1, _ := strconv.Atoi(numbers[2])
		v2, _ := strconv.Atoi(numbers[3])

		robot := Robot{
			Position: Vector{p1, p2},
			Velocity: Vector{v1, v2},
			Time: -1,
		}

		floor[robot.Position.y][robot.Position.x] = append(floor[robot.Position.y][robot.Position.x], robot)

		robots = append(robots, robot)

	}

	file, _ := os.Create("output.txt")
	defer file.Close()

	for i :=1; i<10000;i++ {
		simulate(&floor, i)
		var maybeHasTree bool = checkTreePattern(&floor)
		if (maybeHasTree) {
			draw(&floor, i, file)
		}
	}
}