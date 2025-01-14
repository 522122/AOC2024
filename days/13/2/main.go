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

			currentMachine.Price.X += 10000000000000
			currentMachine.Price.Y += 10000000000000

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

	// an * machine.A.X + bn * machine.B.X = machine.Price.X
	// an * machine.A.Y + bn * machine.B.Y = machine.Price.Y

	bn := (machine.Price.Y * machine.A.X - machine.Price.X * machine.A.Y) / (machine.B.Y * machine.A.X - machine.B.X * machine.A.Y)
	an := (machine.Price.X - bn * machine.B.X) / machine.A.X

	if an * machine.A.X + bn * machine.B.X != machine.Price.X || an * machine.A.Y + bn * machine.B.Y != machine.Price.Y {
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
