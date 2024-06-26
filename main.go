package main

import (
	"fmt"

	"github.com/rcotilla/tic_tac_toe/pkg"
)

func main() {
	board := pkg.NewBoard()

	fmt.Println("Welcome to Tic Tac Toe!")
	fmt.Print("Choose rival: 1. Computer 2. Human: ")
	var rival int
	fmt.Scanf("%d", &rival)

	cmp := pkg.NewComputer("O")

	if rival == 1 {
		fmt.Println("You chose to play against the computer.")
		fmt.Println("Choose mark: 1. X or 2. O:")
		var mark int
		fmt.Scanf("%d", &mark)
		if mark == 1 {
			cmp.Mark = "O"
		} else {
			cmp.Mark = "X"
		}
	} else {
		fmt.Println("You chose to play against a human.")
		fmt.Println("Player 1 will be X and Player 2 will be O.")
	}

	board.PrintBoard()
	if rival == 1 {
		vsComputer(board, cmp)
	} else {
		vsHuman(board)
	}
}

func vsComputer(board *pkg.Board, cmp *pkg.Computer) {
	end_game := false
	msg := ""
	for !end_game {
		var x_or_o = board.GetNext()

		if cmp.Mark == x_or_o {
			cmp.Play(board)
		} else {
			fmt.Println("Enter the position (1-9) for value " + x_or_o + ":")
			var position int
			fmt.Scanf("%d", &position)
			err := makePlay(board, position, x_or_o)
			if err != nil {
				continue
			}
		}
		end_game, msg = board.CheckStatus()
		if end_game {
			fmt.Println(msg)
		}
	}
}

func vsHuman(board *pkg.Board) {
	end_game := false
	msg := ""
	for !end_game {
		var x_or_o = board.GetNext()
		fmt.Println("Enter the position (1-9) for value " + x_or_o + ":")
		var position int
		fmt.Scanf("%d", &position)
		err := makePlay(board, position, x_or_o)
		if err != nil {
			continue
		}
		end_game, msg = board.CheckStatus()
		if end_game {
			fmt.Println(msg)
			fmt.Println("Presione Enter para salir.")
			fmt.Scanln()
		}
	}
}

func makePlay(board *pkg.Board, position int, x_or_o string) error {
	err := board.UpdateBoard(position, x_or_o)
	if err != nil {
		fmt.Println(err)
		_ = board.GetNext()
	}

	return err
}
