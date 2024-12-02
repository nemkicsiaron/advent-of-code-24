package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isReportSafe(report []int) bool {
    ascending, descending, skip := false, false, false

        for j := 1; j < len(report); j++ {
            if report[j-1] == report[j] || report[j-1] - report[j] > 3 || report[j-1] - report[j] < -3 {
                skip = true
                // fmt.Println("SKIPPING in line ", i, " because of",  reports[i][j-1],  reports[i][j])
                break
            }

            descending = descending || report[j-1] > report[j]
            ascending = ascending || report[j-1] < report[j]
        }

        return ((ascending && !descending) || (!ascending && descending)) && !skip
}

func read(path string) [][]int {

    var matrix [][]int

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

        var nums []int
        // Split the line by whitespace and filter out empty strings
        for _, str := range strings.Fields(line) {
            num, err := strconv.Atoi(str)
            if err != nil {
                fmt.Println("Error converting string \"", str, "\"to a number:", err)
            }
            nums = append(nums,num)
        }

        matrix = append(matrix, nums)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading the file:", err)
    }

    return matrix
}

func one(reports [][]int) int {
    totalSafeReports := 0

    for i := 0; i < len(reports); i++ {
        if isReportSafe(reports[i]) {
            totalSafeReports++
        }
    }

    return totalSafeReports
}

func two(reports [][]int) int {
    totalSafeReports := 0

    for i := 0; i < len(reports); i++ {
        fmt.Println("For line nr.", i, " we get: ", reports[i])

        if isReportSafe(reports[i]) {
            totalSafeReports += 1
            //fmt.Println("totalSafeReports = ", totalSafeReports)
        } else {
            for j := range reports[i] {
                micro := make([]int, len(reports[i])-1)
                copy(micro, reports[i][:j])
                copy(micro[j:], reports[i][j+1:])
                //fmt.Println(micro, " with j: ", j, " and ", reports[i])
                if isReportSafe(micro) {
                    fmt.Println("I THINK THIS IS SAFE: ", micro)
                    totalSafeReports += 1
                    break
                }
            }
        }
    }

    return totalSafeReports
}

func main() {
    matrix := read("./input.txt")

    //fmt.Println(matrix)

    fmt.Println("one: ", one(matrix))
    fmt.Println("two: ", two(matrix))
}