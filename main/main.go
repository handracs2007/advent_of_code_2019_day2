// https://adventofcode.com/2019/day/2

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	inputStr := strings.Split("1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,10,1,19,1,6,19,23,2,23,6,27,1,5,27,31,1,31,9,35,2,10,35,39,1,5,39,43,2,43,10,47,1,47,6,51,2,51,6,55,2,55,13,59,2,6,59,63,1,63,5,67,1,6,67,71,2,71,9,75,1,6,75,79,2,13,79,83,1,9,83,87,1,87,13,91,2,91,10,95,1,6,95,99,1,99,13,103,1,13,103,107,2,107,10,111,1,9,111,115,1,115,10,119,1,5,119,123,1,6,123,127,1,10,127,131,1,2,131,135,1,135,10,0,99,2,14,0,0", ",")

	// Convert to numbers
	oriInput := make([]int, len(inputStr))
	input := make([]int, len(inputStr))
	for idx, val := range inputStr {
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("Unable to convert input [%s] to integer. %v.", val, err.Error())
		}

		input[idx] = num
		oriInput[idx] = num
	}

	noun := 0
	verb := 0

	for noun = 0; noun <= 99; noun++ {
		for verb = 0; verb <= 99; verb++ {
			// Reset the input
			copy(input, oriInput)

			input[1] = noun
			input[2] = verb

			// Run the programme
			exit := false
			for i := 0; i < len(input) && !exit; i += 4 {
				command := input[i]
				switch command {
				case 1:
					// Addition command
					num1 := input[input[i+1]]
					num2 := input[input[i+2]]
					res := num1 + num2
					input[input[i+3]] = res

				case 2:
					// Multiplication command
					num1 := input[input[i+1]]
					num2 := input[input[i+2]]
					res := num1 * num2
					input[input[i+3]] = res

				case 99:
					// Exit command
					exit = true
				}
			}

			if input[0] == 19690720 {
				break
			}
		}

		if input[0] == 19690720 {
			break
		}
	}

	fmt.Println(100*noun + verb)
}
