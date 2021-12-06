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
	distance := 0
	depth := 0
	for i := 0 ; i < len(lines); i++ {
		direction := lines[i][0]
		amount, _ := strconv.Atoi(lines[i][1])
		switch direction {
		case "forward": distance += amount
		case "down": depth += amount
		case "up": depth -= amount
		}
	}
	fmt.Println("Day 2a:", distance, depth, distance * depth)

	aim := 0
	distance = 0
	depth = 0
	for i := 0 ; i < len(lines); i++ {
		direction := lines[i][0]
		amount, _ := strconv.Atoi(lines[i][1])
		switch direction {
		case "forward":
			distance += amount
			depth += aim * amount
		case "down": aim += amount
		case "up": aim -= amount
		}
	}
	fmt.Println("Day 2b:", distance, depth, aim, distance * depth)


}

func getLines() [][]string {
	file, err := os.Open("./day2/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, strings.Split(scanner.Text(), " "))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result

}
