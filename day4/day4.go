package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BingoSquare struct {
	number int
	drawn bool
}

type BingoCard struct {
	squares [5][5]BingoSquare
}

func (card *BingoCard) Print() {
	for _, line := range card.squares {
		for _, square := range line {
			if square.drawn {
				fmt.Print("\033[32m")
			} else {
				fmt.Print("\033[31m")
			}
			fmt.Printf("%3d", square.number)
		}
		fmt.Println("\033[0m")
	}
}

func (card *BingoCard) ProcessDraw(draw int) {
	for i := range card.squares {
		for j := range card.squares[i] {
			if card.squares[i][j].number == draw {
				card.squares[i][j].drawn = true
			}
		}
	}
}

func (card *BingoCard) IsWinner() bool {

	for i := 0 ; i<5 ; i++ {
		lineResult := true
		for j := 0 ; j<5 ; j++ {
			lineResult = lineResult && card.squares[i][j].drawn
		}
		if lineResult { return true}
	}
	for i := 0 ; i<5 ; i++ {
		colResult := true
		for j := 0 ; j<5 ; j++ {
			colResult = colResult && card.squares[j][i].drawn
		}
		if colResult { return true}
	}
	return false
}

func (card *BingoCard) Score() int {
	score := 0
	for i := 0 ; i<5 ; i++ {
		for j := 0 ; j<5 ; j++ {
			if !card.squares[i][j].drawn {
				score += card.squares[i][j].number
			}
		}
	}
	return score
}

func main() {
	lines := getLines()

	var bingoCards []*BingoCard
	for i :=1 ; i<len(lines) ; i +=6 {
		card := parseCard(lines[i+1:i+6])
		bingoCards = append(bingoCards, &card)
	}

	for _, draw := range lines[0] {
		haveWinner := false
		for _, card := range bingoCards {
			card.ProcessDraw(draw)
		}
		for _, card := range bingoCards {
			if card.IsWinner() {
				fmt.Printf("Day4a draw[%d] sum[%d] result[%d]\n", draw, card.Score(), draw*card.Score())
				card.Print()
				haveWinner = true
			}
		}
		if haveWinner {break}
	}

	bingoCards = make([]*BingoCard, 0)
	for i :=1 ; i<len(lines) ; i +=6 {
		card := parseCard(lines[i+1:i+6])
		bingoCards = append(bingoCards, &card)
	}

	var winners = map[string]bool{}
	for _, draw := range lines[0] {
		for _, card := range bingoCards {
			card.ProcessDraw(draw)
		}
		for i, card := range bingoCards {
			if card.IsWinner() {
				winners[strconv.Itoa(i)] = true
				if len(winners) == len(bingoCards) {
					fmt.Printf("Day4b draw[%d] sum[%d] result[%d]\n", draw, card.Score(), draw*card.Score())
					card.Print()
					return
				}
			}
		}
	}

}

func parseCard(numbers [][]int) BingoCard {
	result := BingoCard{}
	for i, line := range numbers {
		for j, number := range line {
			result.squares[i][j] = BingoSquare{number: number, drawn: false}
		}
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

func getLines() [][]int {
	file, err := os.Open("./day4/day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := make([]int, 0)
		for _,s := range strings.Split(strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), "  ", " "), " ", ","), ",") {
			if s != "" {
				line = append(line, parseInt(s))
			}
		}
		result = append(result, line)
	}

	return result

}
