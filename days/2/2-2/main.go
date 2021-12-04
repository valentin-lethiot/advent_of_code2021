package main

import (
	"fmt"
	"github.com/valentin-lethiot/advent_of_code2021/days/2/models"
	"github.com/valentin-lethiot/advent_of_code2021/utils"
)

func main() {
	inputsString := utils.GetInputsToString("../inputs.txt")
	inputs := models.StringArrayToInputArray(inputsString)

	distance := 0
	depth := 0
	aim := 0
	for _, input := range inputs {
		switch input.Command {
		case "forward":
			distance += input.Value
			depth = depth + (aim * input.Value)
			break
		case "up":
			aim -= input.Value
			break
		case "down":
			aim += input.Value
			break
		}
	}

	fmt.Println("AIM : ", aim)
	fmt.Println("DISTANCE : ", distance)
	fmt.Println("DEPTH : ", depth)
	fmt.Println("MULTIPLIED : ", distance * depth)
}
