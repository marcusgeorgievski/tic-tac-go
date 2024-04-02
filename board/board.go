package board

import (
	"errors"
	"fmt"

	"github.com/marcusgeorgievski/tictacgo/player"
	"github.com/marcusgeorgievski/tictacgo/position"
	"github.com/marcusgeorgievski/tictacgo/termfmt"
)

type Board [3][3]player.Symbol

var EmptySymbol player.Symbol = player.Symbol(' ')
var ErrSpotTaken = errors.New("spot is already taken")
var ErrIndexOutOfBounds = errors.New("index value is out of bounds [3][3]")

func NewBoard() *Board {
	board := &Board{{}, {}, {}}
	board.EmptyBoard()
	return board
}

func (b *Board) SetSpot(p *position.Position, s player.Symbol) error {
	if b[p.R][p.C] != EmptySymbol {
		return ErrSpotTaken
	}
	b[p.R][p.C] = s
	return nil
}

func (b *Board) Show() {
	for i := range b {
		fmt.Printf(" %v | %v | %v \n", b[i][0], b[i][1], b[i][2])
		if i != 2 {
			divider()
		}
	}
}
func (b *Board) ShowSelection(p *position.Position, col string) {
	highlight := func(i, j int) string {
		if p.R == i && p.C == j {
			return termfmt.SprintfFormat(fmt.Sprint(b[i][j]), termfmt.UNDERLINE, col)
		}
		return fmt.Sprint(b[i][j])
	}

	for i := range b {
		for j := range b[i] {
			if j < 2 {
				fmt.Printf(" %v |", highlight(i, j))
			} else {
				fmt.Printf(" %v \n", highlight(i, j))
			}
		}
		if i != 2 {
			divider()
		}
	}
}

func (b *Board) EmptyBoard() {
	for i := range b {
		for j := range b[i] {
			b[i][j] = player.Symbol(' ')
		}
	}
}

func (b *Board) Erase() {
	termfmt.EraseLine()
	termfmt.EraseLinesAbove(5)
}

func (b *Board) CheckForWin() string {
	for i := range 3 {
		// Check horizontal win
		if b[i][0] == b[i][1] && b[i][0] == b[i][2] && b[i][0] != EmptySymbol {
			return WINNER
		}
		// Check vertical win
		if b[0][i] == b[1][i] && b[0][i] == b[2][i] && b[0][i] != EmptySymbol {
			return WINNER
		}
	}
	// Check top-left to bottom-right win
	if b[0][0] == b[1][1] && b[0][0] == b[2][2] && b[0][0] != EmptySymbol {
		return WINNER
	}

	// Check top-right to bottom-left win
	if b[0][2] == b[1][1] && b[0][0] == b[2][0] && b[0][2] != EmptySymbol {
		return WINNER
	}

	// Check if an empty spot is available
	for i := range 3 {
		for j := range 3 {
			if b[i][j] == EmptySymbol {
				return PLAYING
			}
		}
	}
	return TIE
}

const (
	WINNER  = "WINNER"
	TIE     = "TIE"
	PLAYING = "PLAYING"
)

func divider() {
	fmt.Printf("---+---+---\n")
}
