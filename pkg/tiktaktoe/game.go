package tiktaktoe

import (
	"errors"

	"fyne.io/fyne/v2/widget"
)

type Game struct {
	table [3][3]*widget.Button
	turn  string
}

func NewGame(table [3][3]*widget.Button) *Game {
	for _, row := range table {
		for _, square := range row {
			square.SetText("")
		}
	}
	return &Game{
		table: table,
	}
}
func (g *Game) MakePlay(row int8, column int8) error {
	if g.turn == "" {
		g.turn = "X"
	}
	if row < 0 || row > 2 {
		return errors.New("Invalid row")
	}
	if column < 0 || column > 2 {
		return errors.New("Invalid column")
	}
	square := g.table[row][column]
	if square.Text != "" {
		return errors.New("square already marked")
	}
	square.SetText(g.turn)
	square.Refresh()
	if g.turn == "X" {
		g.turn = "O"
	} else {
		g.turn = "X"
	}
	return nil
}
