package pkg

import (
	"fmt"
	"slices"
	"strings"
)

type Position struct {
	position int
	value    string
}

type Board struct {
	board [9]Position
	value string
}

// NewBoard creates a new board with all cells empty.
func NewBoard() *Board {
	return &Board{
		board: [9]Position{
			{position: 1, value: " "},
			{position: 2, value: " "},
			{position: 3, value: " "},
			{position: 4, value: " "},
			{position: 5, value: " "},
			{position: 6, value: " "},
			{position: 7, value: " "},
			{position: 8, value: " "},
			{position: 9, value: " "},
		},
		value: "X",
	}
}

// PrintBoard prints the current state of the board.
func (b *Board) PrintBoard() {
	separator := "  |  "
	fmt.Println(" " + b.board[0].value + separator + b.board[1].value + separator + b.board[2].value)
	fmt.Println("--- - --- - ---")
	fmt.Println(" " + b.board[3].value + separator + b.board[4].value + separator + b.board[5].value)
	fmt.Println("--- - --- - ---")
	fmt.Println(" " + b.board[6].value + separator + b.board[7].value + separator + b.board[8].value)
}

// UpdateBoard updates the board with the player's move.
func (b *Board) UpdateBoard(position int, value string) {
	possible_values := []string{"X", "O"}
	if !slices.Contains(possible_values, value) {
		fmt.Println("Invalid value. Please enter X or O.")
		return
	} else {
		b.board[position-1].value = value
	}

	fmt.Println(strings.Repeat("+", 15))
	b.PrintBoard()
}

func (b *Board) GetNext() string {
	current_value := b.value
	if b.value == "X" {
		b.value = "O"
	} else {
		b.value = "X"
	}

	return current_value
}

func (b *Board) CheckStatus() (bool, string) {
	has_winner, winner := b.checkWinner()
	if has_winner {
		return true, winner
	}
	return b.checkDraw()
}

func (b *Board) checkDraw() (bool, string) {
	for _, cell := range b.board {
		if cell.value == " " {
			return false, ""
		}
	}

	return true, "Draw"
}

func (b *Board) checkWinner() (bool, string) {
	winning_options := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Rows
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Columns
		{0, 4, 8}, {2, 4, 6}, // Diagonals
	}

	for _, wop := range winning_options {
		if b.isSameValue(wop[0], wop[1], wop[2]) && !b.checkMark(" ", wop[0]) {
			fmt.Println("Player " + b.board[wop[0]].value + " wins!")
			return true, b.board[wop[0]].value
		}
	}

	return false, ""
}

func (b *Board) isSameValue(positions ...int) bool {
	value := b.board[positions[0]].value
	for _, position := range positions {
		if value != b.board[position].value {
			return false
		}
	}
	return true
}

func (b *Board) checkMark(mark string, position int) bool {
	return b.board[position].value == mark
}

func (b *Board) CheckWinnerOption(mark string) int {
	winning_options := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Rows
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Columns
		{0, 4, 8}, {2, 4, 6}, // Diagonals

	}
	wop_check := -1

	for _, wop := range winning_options {
		if b.isSameValue(wop[0], wop[1]) && b.checkMark(" ", wop[2]) && b.checkMark(mark, wop[0]) {
			wop_check = wop[2]
		}
		if b.isSameValue(wop[1], wop[2]) && b.checkMark(" ", wop[0]) && b.checkMark(mark, wop[1]) {
			wop_check = wop[0]
		}
		if b.isSameValue(wop[0], wop[2]) && b.checkMark(" ", wop[1]) && b.checkMark(mark, wop[0]) {
			wop_check = wop[1]
		}
	}

	return wop_check
}
