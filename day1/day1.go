package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines := getLines()
	count1 := 0
	for i := 1 ; i < len(lines); i++ {
		if lines[i] > lines[i-1] { count1++ }
	}
	fmt.Println("Day 1a: ", count1)

	count2 := 0
	for i := 3 ; i < len(lines); i++ {
		if lines[i] + lines[i-1] + lines[i-2] > lines[i-1] + lines[i-2] + lines[i-3] { count2++ }
	}
	fmt.Println("Day 1b: ", count2)
}

func getLines() []int {
	file, err := os.Open("./day1/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		n, err := strconv.ParseInt(scanner.Text(), 10, 32)

		if err != nil {
			log.Fatal("Can't parse to int: " + scanner.Text())
		}

		result = append(result, int(n))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result

}
