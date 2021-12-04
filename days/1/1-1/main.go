package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    inputs := getInputs("../inputs.txt")

    increased := 0
    for i := 1; i < len(inputs); i++ {
        if inputs[i-1] < inputs[i] {
            increased++
        }
    }
    fmt.Println(increased, " augmentations.")
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
