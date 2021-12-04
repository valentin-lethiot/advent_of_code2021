package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    inputs := getInputs("./inputs.txt")

    increased := 0
    for i := 1; i < len(inputs); i++ {
        if inputs[i] > inputs[i-1] {
            increased++
        }
    }
    fmt.Println(increased, " augmentations.")
}

func getInputs(path string) []int {
    inputs := make([]string, 0)

    content, _ := os.Open(path)
    scanner := bufio.NewScanner(content)

    for scanner.Scan() {
        inputs = append(inputs, scanner.Text())
    }

    return inputs
}