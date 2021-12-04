package models

import (
	"strconv"
	"strings"
)

type Input struct {
	Command string
	Value int
}

func NewInput(commandValue string) *Input {
	inputs := strings.Split(commandValue, " ")
	value, _ := strconv.Atoi(inputs[1])

	return &Input{
		Command: inputs[0],
		Value: value,
	}
}

func StringArrayToInputArray(inputs[]string) []*Input {
	returnedInputs := make([]*Input, 0)

	for _, input := range inputs {
		returnedInputs = append(returnedInputs, NewInput(input))
	}

	return returnedInputs
}
