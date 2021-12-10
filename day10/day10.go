package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	lines := getLines()
	resultA := 0
	var resultB []int
	for _, line := range lines {
		scoreA, scoreB := computeScore(line)
		resultA += scoreA
		if scoreB > 0 {
			resultB = append(resultB, scoreB)
		}
	}
	fmt.Printf("Day10a %d\n", resultA)
	sort.Ints(resultB)
	fmt.Printf("Day10b %d\n", resultB[len(resultB)/2])

}

func computeScore(line []string) (int, int) {
	stack := []string{line[0]}
	for i := 1; i < len(line); i++ {
		current := line[i]
		switch current {
		case "(":
			stack = append(stack, current)
		case "[":
			stack = append(stack, current)
		case "{":
			stack = append(stack, current)
		case "<":
			stack = append(stack, current)
		case ")":
			if stack[len(stack)-1] == "(" {
				stack = stack[0:len(stack)-1]
			} else {
				return 3, 0
			}
		case "]":
			if stack[len(stack)-1] == "[" {
				stack = stack[0:len(stack)-1]
			} else {
				return 57, 0
			}
		case "}":
			if stack[len(stack)-1] == "{" {
				stack = stack[0:len(stack)-1]
			} else {
				return 1197, 0
			}
		case ">":
			if stack[len(stack)-1] == "<" {
				stack = stack[0:len(stack)-1]
			} else {
				return 25137, 0
			}
		}
	}
	return 0, scoreStack(stack)
}

func scoreStack(stack []string) int {
	result := 0
	for i := len(stack)-1 ; i >= 0 ; i-- {
		result = result * 5
		switch stack[i] {
		case "(": result += 1
		case "[": result += 2
		case "{": result += 3
		case "<": result += 4
		}
	}
	return result
}


func getLines() [][]string {
	file, err := os.Open("./day10/day10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line []string
		for _, s := range strings.Split(scanner.Text(), "") {
			line = append(line, s)
		}
		result = append(result, line)
	}

	return result
}
