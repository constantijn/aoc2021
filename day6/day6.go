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
	fmt.Println("Day 6a", solve(lines, 80))
	fmt.Println("Day 6a", solve(lines, 256))
}

func solve(lines []int, days int) int {
	counts := map[int]int {0:0, 1:0, 2:0, 3:0, 4:0, 5:0, 6:0, 7:0, 8:0}
	for _,days := range lines{
		counts[days]++
	}

	for i := 1; i <= days; i++ {
		zeroes := counts[0]
		counts[0] = counts[1]
		counts[1] = counts[2]
		counts[2] = counts[3]
		counts[3] = counts[4]
		counts[4] = counts[5]
		counts[5] = counts[6]
		counts[6] = counts[7] + zeroes
		counts[7] = counts[8]
		counts[8] = zeroes
	}

	result := 0
	for _, count := range counts {
		result += count
	}

	return result
}

func parseInt(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("invalid int string [", input, "]")
	}
	return result
}

func getLines() []int {
	file, err := os.Open("./day6/day6.txt")
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
