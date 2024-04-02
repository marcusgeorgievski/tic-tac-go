package main

import (
	"fmt"

	"github.com/marcusgeorgievski/tictacgo/board"
	"github.com/marcusgeorgievski/tictacgo/player"
	"github.com/marcusgeorgievski/tictacgo/position"
	"github.com/marcusgeorgievski/tictacgo/termfmt"

	"github.com/eiannone/keyboard"
)

func main() {
	key := keyboard.Key(0)
	players := [2]player.Player{}
	colors := [2]string{termfmt.RED, termfmt.BLUE}

	termfmt.PrintfFormat("\nWelcome to tic-tac-go!\n\n", termfmt.ORANGE, termfmt.UNDERLINE)

	// Create players
	for i := range 2 {
		termfmt.PrintfFormat("Player "+fmt.Sprint(i+1)+"\n", colors[i])
		name, symbol := termfmt.GetPlayerInfo()
		players[i].Name = name
		players[i].Symbol = player.Symbol(symbol[0])
	}

	pos := position.Position{} // tracks what cell player is hovering on
	b := board.Board{}         // board
	b.EmptyBoard()

	termfmt.PrintfFormat("Begin!\n\n", termfmt.RED)

	play := true
	// Main game loop
	for turn := 1; play; turn++ {

		termfmt.EraseLinesAbove(2)

		if turn%2 == 0 {
			termfmt.PrintfFormat("Player 2's turn\n\n", colors[1])
		} else {
			termfmt.PrintfFormat("Player 1's turn\n\n", colors[0])
		}

		// Loops while player is moving in grid
		for moved := false; !moved; {
			b.ShowSelection(&pos, colors[currPlayer(turn)])
			key = termfmt.GetKey()
			b.Erase()
			switch key {
			case termfmt.UP:
				pos.Up()
			case termfmt.DOWN:
				pos.Down()
			case termfmt.LEFT:
				pos.Left()
			case termfmt.RIGHT:
				pos.Right()
			case termfmt.ENTER:
				var err error
				if turn%2 == 0 {
					err = b.SetSpot(&pos, players[1].Symbol)
				} else {
					err = b.SetSpot(&pos, players[0].Symbol)
				}
				if err == board.ErrSpotTaken {
					turn-- // stay on same turn
				}
				moved = true
			case termfmt.QUIT:
				play = false
				moved = true
			}

		}

		result := b.CheckForWin()
		if result == board.WINNER {
			termfmt.EraseLinesAbove(2)
			b.Show()
			if turn%2 == 0 {
				termfmt.PrintfFormat("\nPlayer "+fmt.Sprint(2)+" won!\n", colors[1])
			} else {
				termfmt.PrintfFormat("\nPlayer "+fmt.Sprint(1)+" won!\n", colors[0])
			}
			play = false
		} else if result == board.TIE {
			termfmt.EraseLinesAbove(2)
			termfmt.PrintfFormat("Tie!\n", termfmt.ORANGE)
			play = false
		}
	}

	fmt.Printf("\nThanks for playing!\n")
	termfmt.PrintfFormat("https://github.com/marcusgeorgievski\n", termfmt.GREEN)
}

func currPlayer(turn int) int {
	if turn%2 == 0 {
		return 1
	}
	return 0
}
