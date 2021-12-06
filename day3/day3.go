package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := getLines()

	bitCount := len(lines[0])
	gamma := ""
	epsilon := ""

	for i := 0 ; i < bitCount; i++ {
		zeros := 0
		ones := 0
		for j := 0 ; j < len(lines); j++ {
			switch lines[j][i] {
			case 0: zeros++
			case 1: ones++
			}
		}
		if zeros > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaInt,_ := strconv.ParseInt(gamma, 2, bitCount+1)
	epsilonInt,_ := strconv.ParseInt(epsilon, 2, bitCount+1)

	fmt.Println("Day 3a", gammaInt, epsilonInt, gammaInt * epsilonInt)

	oxygen := part2(bitCount, lines, 1)
	co2 := part2(bitCount, lines, 0)

	fmt.Println("Day 3b", oxygen, co2, oxygen * co2)
}

func part2(bitCount int, lines [][]int, filterOn int) int64 {
	for i := 0; i < bitCount; i++ {
		zeros := 0
		ones := 0
		var newLines [][]int
		for j := 0; j < len(lines); j++ {
			switch lines[j][i] {
			case 0:
				zeros++
			case 1:
				ones++
			}
		}
		if zeros > ones {
			for j := 0; j < len(lines); j++ {
				if lines[j][i] != filterOn {
					newLines = append(newLines, lines[j])
				}
			}
		} else {
			for j := 0; j < len(lines); j++ {
				if lines[j][i] == filterOn {
					newLines = append(newLines, lines[j])
				}
			}
		}
		lines = newLines

		if len(lines) == 1 {
			break
		}
	}

	resultString := ""
	for _, s := range lines[0] {
		resultString += strconv.Itoa(s)
	}

	resultInt,_ := strconv.ParseInt(resultString, 2, bitCount+1)


	return resultInt
}

func parseInt(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("invalid int string [", input, "]")
	}
	return result
}

func getLines() [][]int {
	file, err := os.Open("./day3/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var intLine []int
		for _, s := range strings.Split(scanner.Text(), "") {
			intLine = append(intLine, parseInt(s))
		}
		result = append(result, intLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result

}
