package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
)

const (
	costA = 3
	costB = 1
)

type Vector struct {
	X int
	Y int
}

type ClawMachine struct {
	A Vector
	B Vector
	Price Vector
}

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseInput(input *[]string) []ClawMachine {
	re := regexp.MustCompile(`[-]?\d+`)
	machines := []ClawMachine{}
	currentMachine := ClawMachine{}

	skipped := 0
	for i, line := range *input {
		if line == "" {
			skipped++
			continue
		}
		matches := re.FindAllString(line, -1)
		switch (i + 1 - skipped) % 3 {
		case 0:
			// fmt.Println("Price", matches)

			currentMachine.Price.X, _ = strconv.Atoi(matches[0])
			currentMachine.Price.Y, _ = strconv.Atoi(matches[1])

			machines = append(machines, currentMachine)
			currentMachine = ClawMachine{}
		case 1:
			// fmt.Println("A", matches)

			currentMachine.A.X, _ = strconv.Atoi(matches[0])
			currentMachine.A.Y, _ = strconv.Atoi(matches[1])
		case 2:
			// fmt.Println("B", matches)

			currentMachine.B.X, _ = strconv.Atoi(matches[0])
			currentMachine.B.Y, _ = strconv.Atoi(matches[1])
		}
	}

	return machines
}

func cost(an int, bn int) int {
	return (an * costA) + (bn * costB)
}

func solveMachine(machine ClawMachine) (int, int, int) {
	an := 0
	bn := 0
	maxB := false
	current := Vector{X: 0, Y: 0}

	for {
		if current.X + machine.B.X <= machine.Price.X && current.Y + machine.B.Y <= machine.Price.Y && !maxB {
			current.X += machine.B.X
			current.Y += machine.B.Y
			bn++
		} else if current.X + machine.A.X <= machine.Price.X && current.Y + machine.A.Y <= machine.Price.Y {
			current.X += machine.A.X
			current.Y += machine.A.Y
			an++
		} else if current.X == machine.Price.X && current.Y == machine.Price.Y {
			break
		} else {
			maxB = true
			current.X -= machine.B.X
			current.Y -= machine.B.Y
			bn--
			if bn < 0 {
				return 0, 0, 0
			}
			// fmt.Println("Correction")
		}
	}

	if an > 100 || bn > 100 {
		return 0, 0, 0
	}

	return an, bn, cost(an, bn)
}

func main() {
	var input []string = readInput()

	var machines []ClawMachine = parseInput(&input)

	sum := 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for _, machine := range machines {
		wg.Add(1)
		go func(machine ClawMachine) {
			defer wg.Done()
			fmt.Println("Solving machine", machine)
			an, bn, cost := solveMachine(machine)
			fmt.Println(an, bn, cost)
			if cost > 0 {
				mu.Lock()
				sum += cost
				mu.Unlock()
			}
		} (machine)
	}

	wg.Wait()
	fmt.Println(sum)

}
