package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Cell struct {
	height int
	basin int
}

func main() {
	grid := getLines()
	count := 0
	for i := 0 ; i < len(grid) ; i++ {
		for j := 0; j < len(grid[i]); j++ {
			count += isLowPoint(i,j,grid)
		}
	}
	fmt.Printf("Day9a %d\n", count)
	var cells [][]Cell
	for i := 0 ; i < len(grid) ; i++ {
		var line []Cell
		for j := 0; j < len(grid[i]); j++ {
			line = append(line, Cell{height: grid[i][j]})
		}
		cells = append(cells, line)
	}

	sizeByBasin := make(map[int]int)
	basinId := 0
	for i := 0 ; i < len(grid) ; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if cells[i][j].height != 9 && cells[i][j].basin == 0 {
				basinId++
				sizeByBasin[basinId] = 0
				fillBasin(basinId, i, j, cells, sizeByBasin)
			}
		}
	}

	//printCells(cells)
	//fmt.Printf("BasinSizes: %+v\n", sizeByBasin)

	basinSizes := make([]int, 0, len(sizeByBasin))
	for  _, value := range sizeByBasin {
		basinSizes = append(basinSizes, value)
	}
	sort.Ints(basinSizes)
	//fmt.Printf("BasinSizes: %+v\n", basinSizes)
	fmt.Printf("Day9b %d\n", basinSizes[len(basinSizes) - 1] * basinSizes[len(basinSizes) - 2] * basinSizes[len(basinSizes) - 3])

}

func fillBasin(basin int, i int, j int, cells [][]Cell, sizeByBasin map[int]int) {
	//fmt.Printf("fillBasin %d,%d -> %d\n", i, j, basin)
	cells[i][j].basin = basin
	sizeByBasin[basin]++
	if i > 0 && cells[i-1][j].height != 9 && cells[i-1][j].basin == 0 {
		fillBasin(basin, i - 1, j, cells, sizeByBasin)
	}
	if i < len(cells) - 1 && cells[i+1][j].height != 9 && cells[i+1][j].basin == 0{
		fillBasin(basin, i + 1, j, cells, sizeByBasin)
	}
	if j > 0 && cells[i][j-1].height != 9 && cells[i][j-1].basin == 0{
		fillBasin(basin, i, j - 1, cells, sizeByBasin)
	}
	if j < len(cells[i]) - 1 && cells[i][j+1].height != 9 && cells[i][j+1].basin == 0 {
		fillBasin(basin, i, j + 1, cells, sizeByBasin)
	}
}

func printCells(cells [][]Cell) {
	for i := 0 ; i < len(cells) ; i++ {
		for j := 0; j < len(cells[i]); j++ {
			print(cells[i][j].basin)
		}
		println()
	}
}

func isLowPoint(i int, j int, grid [][]int) int {
	cell := grid[i][j]
	if i > 0 && cell >= grid[i-1][j] {
		return 0
	}
	if i < len(grid) - 1 && cell >= grid[i+1][j] {
		return 0
	}
	if j > 0 && cell >= grid[i][j-1] {
		return 0
	}
	if j < len(grid[i]) - 1 && cell >= grid[i][j+1] {
		return 0
	}
	//fmt.Printf("%d,%d found\n", i,j)
	return cell + 1
}

func parseInt(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("invalid int string [", input, "]")
	}
	return result
}

func getLines() [][]int {
	file, err := os.Open("./day9/day9.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line []int
		for _, s := range strings.Split(scanner.Text(), "") {
			line = append(line, parseInt(s))
		}
		result = append(result, line)
	}

	return result
}
