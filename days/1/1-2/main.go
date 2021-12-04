package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    inputs := getInputs("../inputs.txt")

    sums := getSumOfThree(inputs)

    increased := 0
    for i := 1; i < len(sums); i++ {
        if sums[i-1] < sums[i] {
            increased++
        }
    }
    fmt.Println(increased, " augmentations.")
}

func getSumOfThree(inputs []int) []int {
    sums := make([]int, 0)

    for i := 2; i < len(inputs); i++ {
        sum := inputs[i] + inputs[i-1] + inputs[i-2]
        sums = append(sums, sum)
    }

    return sums
}

func getInputs(path string) []int {
    inputs := make([]int, 0)

    content, _ := os.Open(path)
    scanner := bufio.NewScanner(content)

    for scanner.Scan() {
        input, _ := strconv.Atoi(scanner.Text())
        inputs = append(inputs, input)
    }

    return inputs
}
