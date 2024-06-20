package pkg

type Computer struct {
	Mark string
}

// NewComputer creates a new computer with the given mark.
func NewComputer(mark string) *Computer {
	return &Computer{
		Mark: mark,
	}
}

// Play makes a move on the board.
func (c *Computer) Play(b *Board) {
	position := c.findBestMove(b)
	b.UpdateBoard(position, c.Mark)
}

func (c *Computer) getOponentMark() string {
	if c.Mark == "X" {
		return "O"
	}

	return "X"
}

func (c *Computer) findBestMove(b *Board) int {
	// Check if the computer can win in the next move
	winnerOption := b.CheckWinnerOption(c.Mark)
	if winnerOption > 0 {
		return winnerOption
	}

	// Check if the player can win in the next move
	oponentMark := c.getOponentMark()
	avoidOponentWinOption := b.CheckWinnerOption(oponentMark)
	if avoidOponentWinOption > 0 {
		return avoidOponentWinOption + 1
	}

	// Check if the center is available
	if b.board[4].value == " " {
		return 5
	}

	// Check if a corner is available
	corners := []int{0, 2, 6, 8}
	for _, corner := range corners {
		if b.board[corner].value == " " {
			return corner + 1
		}
	}

	// Check if a side is available
	sides := []int{1, 3, 5, 7}
	for _, side := range sides {
		if b.board[side].value == " " {
			return side + 1
		}
	}

	return 0
}
