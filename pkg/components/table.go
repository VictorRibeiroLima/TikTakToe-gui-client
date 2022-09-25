package components

import (
	"tiktaktoe/pkg/tiktaktoe"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
)

type Table struct {
	game         *tiktaktoe.Game
	c            *websocket.Conn
	roomLabel    *widget.Label
	yourTurn     bool
	messageLabel *widget.Label
	resultLabel  *widget.Label
	table        [3][3]*widget.Button
	gameStarted  bool
	grid         *fyne.Container
	Stop         bool
}

func (t *Table) CreateTable(c *websocket.Conn) {
	b1 := widget.NewButton("", func() {

		t.makePlay(0, 0)

	})
	b2 := widget.NewButton("", func() {

		t.makePlay(0, 1)

	})
	b3 := widget.NewButton("", func() {

		t.makePlay(0, 2)

	})
	b4 := widget.NewButton("", func() {

		t.makePlay(1, 0)

	})
	b5 := widget.NewButton("", func() {

		t.makePlay(1, 1)

	})
	b6 := widget.NewButton("", func() {

		t.makePlay(1, 2)

	})
	b7 := widget.NewButton("", func() {

		t.makePlay(2, 0)

	})
	b8 := widget.NewButton("", func() {
		t.makePlay(2, 1)

	})
	b9 := widget.NewButton("", func() {

		t.makePlay(2, 2)

	})

	table := [3][3]*widget.Button{{b1, b2, b3}, {b4, b5, b6}, {b7, b8, b9}}
	grid := container.New(layout.NewGridLayout(3), b1, b2, b3, b4, b5, b6, b7, b8, b9)
	t.table = table
	t.grid = grid
	t.roomLabel = roomLabel
	t.messageLabel = messageLabel
	t.resultLabel = resultLabel
	t.c = c
}

func (t *Table) makePlay(row int8, column int8) {
	if t.gameStarted && t.yourTurn {
		err := t.game.MakePlay(row, column)
		if err == nil {
			movement := tiktaktoe.Movement{Row: row, Column: column}
			t.c.WriteJSON(tiktaktoe.Message{Event: "PLAY", Message: "Movement", Movement: movement})
		}
		t.yourTurn = false
	}
}

func (t *Table) GameLoop() {
	var message tiktaktoe.Message
	for !t.Stop {
		err := t.c.ReadJSON(&message)
		if err != nil {
			return
		}
		event := message.Event
		if event == "ERROR" {
			t.messageLabel.SetText("Message:" + message.Message)
		} else if event == "ROOM_CONNECTION" {
			t.roomLabel.SetText(message.Message)
			t.messageLabel.SetText("Waiting other player")
		} else if event == "GAME_START" {
			t.game = tiktaktoe.NewGame(t.table)
			t.gameStarted = true
		} else if event == "YOUR_TURN" {
			t.yourTurn = true
			t.messageLabel.SetText("Message:" + message.Message)
		} else if event == "OPPONENT_TURN" {
			t.yourTurn = false
			t.messageLabel.SetText("Message:" + message.Message)
		} else if event == "MOVEMENT" {
			t.game.MakePlay(message.Movement.Row, message.Movement.Column)
		} else if event == "WIN" || event == "LOST" || event == "DRAW" {
			t.resultLabel.SetText("Message:" + message.Message)
		} else if event == "GAME_STOP" {
			t.messageLabel.SetText("Message:" + message.Message)
			t.gameStarted = false
		}
	}
}

func (t Table) Grid() *fyne.Container {
	return t.grid
}
