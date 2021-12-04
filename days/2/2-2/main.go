package main

import (
	"github.com/valentin-lethiot/advent_of_code2021/utils"
	"bufio"
	"fmt"
	"os"
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
	inputsString := GetInputsToString("../inputs.txt")
	inputs := InputStringArrayToInputObjectArray(inputsString)

	distance := 0
	depth := 0
	aim := 0
	for _, input := range inputs {
		switch input.command {
		case "forward":
			distance += input.value
			depth = depth + (aim * input.value)
			break
		case "up":
			aim -= input.value
			break
		case "down":
			aim += input.value
			break
		}
	}

	fmt.Println("AIM : ", aim)
	fmt.Println("DISTANCE : ", distance)
	fmt.Println("DEPTH : ", depth)
	fmt.Println("MULTIPLIED : ", distance * depth)
}


func GetInputsToString(path string) []string {
	inputs := make([]string, 0)

	content, _ := os.Open(path)
	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	return inputs
}

func GetInputsToInt(path string) []int {
	inputsString := GetInputsToString(path)

	inputs := make([]int, 0)
	for _, s := range inputsString {
		intFromString, _ := strconv.Atoi(s)
		inputs = append(inputs, intFromString)
	}

	return inputs
}
