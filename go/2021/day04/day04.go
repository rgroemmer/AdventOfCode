package main

import (
	"adventOfCode/helper"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	Fields [5][5]string
}

func newBoard(rows []string) Board {
	var fields = [5][5]string{}
	for i := 0; i < 5; i++ {
		row := strings.Split(rows[i], " ")
		j := 0
		for _, value := range row {
			if value == "" {
				continue
			}
			fields[i][j] = value
			j++
		}
	}
	return Board{Fields: fields}
}

func containsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsBoard(b []Board, e Board) bool {
	for _, a := range b {
		if a == e {
			return true
		}
	}
	return false
}

func (b *Board) checkWin(playedNumbers []string) bool {
	for j := 0; j < 5; j++ {
		horizontal := b.Fields[0][j]
		vertical := b.Fields[j][0]

		//check vertical rows
		if containsString(playedNumbers, horizontal) {
			for i := 1; i < 5; i++ {
				value := b.Fields[i][j]
				if !containsString(playedNumbers, value) {
					break
				}
				if i == 4 {
					return true
				}
			}
		}

		//check horizontal rows
		if containsString(playedNumbers, vertical) {
			for i := 1; i < 5; i++ {
				value := b.Fields[j][i]
				if !containsString(playedNumbers, value) {
					break
				}
				if i == 4 {
					return true
				}
			}
		}
	}

	return false
}

func (b *Board) getAll() []string {
	var all []string
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			all = append(all, b.Fields[i][j])
		}
	}
	return all
}

func (b *Board) print() {
	fmt.Println("\n#-----#-----#-----#-----#-----#")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			value := b.Fields[i][j]
			if len(value) == 1 {
				value = "0" + value
			}
			fmt.Print(" [" + value + "] ")
		}
		println("")
	}
	fmt.Println("#-----#-----#-----#-----#-----#")
}

func calcBoardScore(board Board, playedNumbers []string) int {
	var score int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			value := board.Fields[i][j]
			if !containsString(playedNumbers, value) {
				numb, err := strconv.Atoi(value)
				if err != nil {
					fmt.Println(err)
				}
				score += numb
			}
		}
	}
	return score
}

func getBoardsAndNumbers(input []string) (playingNumbers []string, boards []Board) {
	var fields []string

	for i, row := range input {
		if i == 0 {
			playingNumbers = strings.Split(row, ",")
			continue
		}
		if row != "" && len(fields) < 5 {
			fields = append(fields, row)
		}
		if len(fields) == 5 {
			boards = append(boards, newBoard(fields))
			fields = nil
		}
	}
	return
}

func part1(input []string) (result int) {
	numbers2Play, boards := getBoardsAndNumbers(input)
	var numbersPlayed []string
	var winner Board
	var winningNumber int
	var hasWon bool

	for round := 0; !hasWon; round++ {
		number := numbers2Play[round]
		numbersPlayed = append(numbersPlayed, number)
		if len(numbersPlayed) > 4 {
			for _, board := range boards {
				if board.checkWin(numbersPlayed) {
					winner = board
					winningNumber, _ = strconv.Atoi(number)
					hasWon = true
					break
				}
			}
		}
	}
	score := calcBoardScore(winner, numbersPlayed)

	return winningNumber * score
}

func part2(input []string) (result int) {
	numbers2Play, boards := getBoardsAndNumbers(input)
	var numbersPlayed []string
	var winners []Board
	var winningNumber int
	var isLast bool

	for round := 0; !isLast; round++ {
		number := numbers2Play[round]
		numbersPlayed = append(numbersPlayed, number)
		if len(numbersPlayed) > 4 {
			for _, board := range boards {
				if board.checkWin(numbersPlayed) {
					if !containsBoard(winners, board) {
						winners = append(winners, board)
					}
					if len(winners) == len(boards) {
						winningNumber, _ = strconv.Atoi(number)
						isLast = true
						break
					}
				}
			}
		}
	}

	lastWinner := winners[len(winners)-1]
	score := calcBoardScore(lastWinner, numbersPlayed)

	return winningNumber * score
}

func main() {
	r, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input, err := helper.InputToStringSlice(string(r))
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
