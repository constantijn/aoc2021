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

	count := 0
	for _, line := range lines {
		if is1478(line[10]) {
			count++
		}
		if is1478(line[11]) {
			count++
		}
		if is1478(line[12]) {
			count++
		}
		if is1478(line[13]) {
			count++
		}
	}
	fmt.Printf("Day8a %d\n", count)

	total := 0
	for _, line := range lines {
		total += decode(line)
	}
	fmt.Printf("Day8b %d\n", total)
}

func is1478(input string) bool {
	l := len(input)
	return l == 2 || l == 3 || l == 4 || l == 7
}

func decode(input []string) int {
	strToInt := map[string]int {}
	intToStr := map[int]string {}

	for i := 0 ; i < 10 ; i++ {
		switch len(input[i]) {
		case 2:
			strToInt[input[i]] = 1
			intToStr[1] = input[i]
		case 3:
			strToInt[input[i]] = 7
			intToStr[7] = input[i]
		case 4:
			strToInt[input[i]] = 4
			intToStr[4] = input[i]
		case 7:
			strToInt[input[i]] = 8
			intToStr[8] = input[i]
		}
	}

	for i := 0 ; i < 10 ; i++ {
		if len(input[i]) == 6 {
			diff := difference(input[i], intToStr[1])
			if len(diff) == 5 {
				strToInt[input[i]] = 6
				intToStr[6] = input[i]
			}
		}
	}

	for i := 0 ; i < 10 ; i++ {
		if len(input[i]) == 5 {
			diff := difference(input[i], intToStr[6])
			if len(diff) == 0 {
				strToInt[input[i]] = 5
				intToStr[5] = input[i]
			}
		}
	}

	for i := 0 ; i < 10 ; i++ {
		if len(input[i]) == 6 && input[i] != intToStr[6]{
			diff := difference(input[i], intToStr[5])
			switch len(diff) {
			case 1:
				strToInt[input[i]] = 9
				intToStr[9] = input[i]
			case 2:
				strToInt[input[i]] = 0
				intToStr[0] = input[i]
			}
		}
	}

	for i := 0 ; i < 10 ; i++ {
		if len(input[i]) == 5 && input[i] != intToStr[5]{
			diff := difference(input[i], intToStr[9])
			switch len(diff) {
			case 0:
				strToInt[input[i]] = 3
				intToStr[3] = input[i]
			case 1:
				strToInt[input[i]] = 2
				intToStr[2] = input[i]
			}
		}
	}

	return 1000 * strToInt[input[10]] + 100 * strToInt[input[11]] + 10 * strToInt[input[12]] + strToInt[input[13]]
}

func difference(a, b string) []string {
	mb := make(map[rune]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, string(x))
		}
	}
	return diff
}


func getLines() [][]string {
	file, err := os.Open("./day8/day8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line []string
		for _, s := range strings.Split(strings.Replace(scanner.Text(), "| ", "", 1), " ") {
			parts := strings.Split(s, "")
			sort.Strings(parts)
			line = append(line, strings.Join(parts, ""))
		}
		result = append(result, line)
	}
	return result
}
