package utils

import (
	"bufio"
	"os"
	"strconv"
)

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
