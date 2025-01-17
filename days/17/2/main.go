package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseInput(input *[]string) (int, int, int, []uint8) {
	var instructions []uint8

	sA := strings.Split((*input)[0], "Register A: ")[1]
	sB := strings.Split((*input)[1], "Register B: ")[1]
	sC := strings.Split((*input)[2], "Register C: ")[1]

	sInstructions := strings.Split(strings.Split((*input)[4], "Program: ")[1], ",")

	A, _ := strconv.Atoi(sA)
	B, _ := strconv.Atoi(sB)
	C, _ := strconv.Atoi(sC)

	for _, s := range sInstructions {
		i, _ := strconv.Atoi(s)
		instructions = append(instructions, uint8(i))
	}

	return A, B, C, instructions
}

func getComboOperand(operand uint8, A *int, B *int, C *int) (int, error) {
	switch operand {
	case 0, 1, 2, 3:
		return int(operand), nil
	case 4:
		return *A, nil
	case 5:
		return *B, nil
	case 6:
		return *C, nil
	default:
		return -1, fmt.Errorf("invalid operand: %d", operand)
	}
}

func runInstruction(A *int, B *int, C *int, instructions []uint8, pointer *int) (int, error) {
	movePointerBy := 2
	output := -1

	if *pointer > len(instructions)-2 {
		return output, fmt.Errorf("pointer out of bounds: %d", *pointer)
	}

	opcode := instructions[*pointer]
	operand := instructions[*pointer+1]

	comboOperand, _ := getComboOperand(operand, A, B, C)

	// fmt.Println("opcode", opcode, "operand", operand, "comboOperand", comboOperand)

	switch opcode {
	case 0: // adv
		// division numerator is A, denominator is 2 to power of comboOperand
		// truncated to integer and witten to A
		*A = int(*A / (1 << comboOperand))
	case 1: // bxl
		// bitwise XOR of B and operand, written to B
		*B = *B ^ int(operand)
	case 2: // bst
		// combo operand modulo 8 writes to B
		*B = comboOperand % 8
	case 3: // jnz
		// nothing if A is 0
		// otherwise, jump pointer to the value of literal, not increasing pointer by 2
		if *A != 0 {
			*pointer = int(operand)
			movePointerBy = 0
		}
	case 4: // bxc
		// bitwise XOR of B and C, written to B
		// ignores operand
		*B = *B ^ *C
	case 5: // out
		// value of comboOperand modulo 8, outputs the value
		output = comboOperand % 8
	case 6: //bdv
		// as adv but result is stored in B (still reads from A)
		*B = int(*A / (1 << comboOperand))
	case 7: // cdv
		// exactly like adv but result is stored in C
		*C = int(*A / (1 << comboOperand))
	}

	*pointer += movePointerBy

	return output, nil
}

func runProgram(A int, B int, C int, instructions []uint8) []uint8 {

	pointer := 0
	programOutput := []uint8{}
	for {
		out, error := runInstruction(&A, &B, &C, instructions, &pointer)
		if error != nil {
			break
		}

		if out != -1 {
			programOutput = append(programOutput, uint8(out))
		}
	}

	return programOutput
}

func bruteForce(B int, C int, instructions []uint8, precession int) int {
	A := 1
	Max := -1
	n := 1
	for {
		programOutput := runProgram(A, B, C, instructions)

		// until we start producing the same length
		if Max == -1 {
			if len(programOutput) == len(instructions) {
				fmt.Println("Top found", A, programOutput, len(programOutput), len(instructions))
				Max = A
				continue
			} else {
				A *= 10
			}
		} else {
			if reflect.DeepEqual(programOutput[len(programOutput)-n:], instructions[len(instructions)-n:]) {
				if n == len(instructions) {
					fmt.Println("Perfect match", A, programOutput, len(programOutput), len(instructions))
					return A
				} else {
					fmt.Println("Found", A, programOutput, len(programOutput), len(instructions))
					n += 1
				}
			}

			A += max(Max/int(math.Pow(10, float64(n+precession))), 1)
		}
	}
}

func main() {
	var input []string = readInput()
	_, B, C, instructions := parseInput(&input)

	A := bruteForce(B, C, instructions, 2)

	fmt.Println(A)

}
