package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func read(path string) ([]int, []int) {

    var lhs, rhs []int

    // Open the text file
    file, err := os.Open(path) // Replace "input.txt" with your file name
    if err != nil {
        fmt.Println("Error opening the file:", err)
        panic(err)
    }
    defer file.Close()

    // Create a scanner to read the file line by line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        // Split the line by whitespace and filter out empty strings
        parts := strings.Fields(line)
        if len(parts) < 2 {
            fmt.Println("Line does not contain enough numbers:", line)
            continue
        }

        // Convert the first and second parts to integers
        num1, err1 := strconv.Atoi(parts[0])
        num2, err2 := strconv.Atoi(parts[1])

        if err1 != nil || err2 != nil {
            fmt.Println("Error converting to integer:", line)
            continue
        }

        lhs = append(lhs, num1)
        rhs = append(rhs, num2)

        // Process the numbers (for example, print them)
        // fmt.Printf("Number 1: %d, Number 2: %d\n", num1, num2)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading the file:", err)
    }

    return lhs, rhs
}

func main() {
    first, second := read("./input.txt")

    //fmt.Println(first)
    //fmt.Println(second)

    totalDistance := 0

    sort.Ints(first)
    sort.Ints(second)

    for i := 0; i < len(first); i++ {
        if first[i] >= second[i] {
            totalDistance = totalDistance + (first[i] - second[i])
        } else {
            totalDistance = totalDistance + (second[i] - first[i])
        }
    }

    fmt.Println(totalDistance)

    similarities := make(map[int]int)
    numerosity := make(map[int]int)

    for i := 0; i < len(first); i++ {
        similarities[first[i]] = 0
        numerosity[first[i]] = numerosity[first[i]] + 1
    }

    for i := 0; i < len(second); i++ {
        if _, exists := similarities[second[i]]; exists {
            similarities[second[i]] = similarities[second[i]] + 1
        }
    }

    similarityScore := 0

    for k, v := range similarities {
        similarityScore = similarityScore + (k * v * numerosity[k])
    }

    fmt.Println(similarities)
    fmt.Println(numerosity)
    fmt.Println(similarityScore)
}