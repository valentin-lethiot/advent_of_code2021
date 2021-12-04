package main

import (
	"github.com/valentin-lethiot/advent_of_code2021/utils"
	"fmt"
	"strconv"
	"strings"
)

type Input struct {
	command string
	value int
}

func NewInput(commandValue string) *Input {
	inputs := strings.Split(commandValue, " ")
	value, _ := strconv.Atoi(inputs[1])

	return &Input{
		command: inputs[0],
		value: value,
	}
}

func InputStringArrayToInputObjectArray(inputs[]string) []*Input {
	returnedInputs := make([]*Input, 0)

	for _, input := range inputs {
		returnedInputs = append(returnedInputs, NewInput(input))
	}

	return returnedInputs
}

func main() {
	inputsString := utils.GetInputsToString("../inputs.txt")
	inputs := InputStringArrayToInputObjectArray(inputsString)

	distance := 0
	depth := 0
	for _, input := range inputs {
		switch input.command {
		case "forward":
			distance += input.value
			break
		case "up":
			depth -= input.value
			break
		case "down":
			depth += input.value
			break
		}
	}

	fmt.Println("DISTANCE : ", distance)
	fmt.Println("DEPTH : ", depth)
	fmt.Println("MULTIPLIED : ", distance * depth)


}
