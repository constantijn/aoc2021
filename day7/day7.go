package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := getLines()

	min := math.MaxInt
	max := 0
	for _, depth := range lines {
		if depth < min {
			min = depth
		}
		if depth > max {
			 max = depth
		}
	}
	//fmt.Printf("%d %d %v\n", min, max, lines)

	minFuel := math.MaxInt
	for i := min ; i <= max ; i++ {
		fuel := 0
		for _, depth := range lines {
			fuel += abs(depth - i)
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	fmt.Printf("Day7a %d\n", minFuel)

	minFuel = math.MaxInt
	for i := min ; i <= max ; i++ {
		fuel := 0
		for _, depth := range lines {
			n := abs(depth - i)
			fuel += (n * (n + 1)) / 2
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	fmt.Printf("Day7b %d\n", minFuel)


}

func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func parseInt(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("invalid int string [", input, "]")
	}
	return result
}

func getLines() []int {
	file, err := os.Open("./day7/day7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, s := range strings.Split(scanner.Text(), ",") {
			result = append(result, parseInt(s))
		}
	}
	return result
}
