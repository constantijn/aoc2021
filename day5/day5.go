package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const gridSize = 1000

func main() {
	fmt.Println("Day 5a", solve(false))
	fmt.Println("Day 5b", solve(true))
}

func solve(diagonals bool) int {
	var grid [gridSize][gridSize]int

	for _, line := range getLines() {
		if line[0] == line[2] {
			y := line[0]
			xMin := min(line[1], line[3])
			xMax := max(line[1], line[3])
			for i := xMin; i <= xMax; i++ {
				grid[i][y]++
			}
		} else if line[1] == line[3] {
			x := line[1]
			yMin := min(line[0], line[2])
			yMax := max(line[0], line[2])
			for i := yMin; i <= yMax; i++ {
				grid[x][i]++
			}
		} else if diagonals {
			if line[0] < line[2] && line[1] < line[3] {
				for i := 0 ; i <= line[2] - line[0] ; i++ {
					grid[line[1]+i][line[0]+i]++
				}
			} else if line[0] < line[2] && line[1] > line[3] {
				for i := 0 ; i <= line[2] - line[0] ; i++ {
					grid[line[1]-i][line[0]+i]++
				}
			} else if line[0] > line[2] && line[1] < line[3] {
				for i := 0 ; i <= line[0] - line[2] ; i++ {
					grid[line[1]+i][line[0]-i]++
				}
			} else if line[0] > line[2] && line[1] > line[3] {
				for i := 0 ; i <= line[0] - line[2] ; i++ {
					grid[line[1]-i][line[0]-i]++
				}
			}

		}
	}

	//printGrid(grid)

	count := 0
	for _, line := range grid {
		for _, point := range line {
			if point >= 2 {
				count++
			}
		}
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printGrid(grid [gridSize][gridSize]int) {
	for _, line := range grid {
		for _, point := range line {
			if point == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%d", point)
			}
		}
		fmt.Print("\n")
	}
}

func parseInt(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("invalid int string [", input, "]")
	}
	return result
}

func getLines() [][]int {
	file, err := os.Open("./day5/day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line []int
		for _, s := range strings.Split(strings.Replace(scanner.Text(), " -> ", ",", 1), ",") {
			line = append(line, parseInt(s))
		}
		result = append(result, line)
	}

	return result

}
