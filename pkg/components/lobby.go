package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
)

func CreateLobby(window fyne.Window, c *websocket.Conn, quitTo func()) {
	table := Table{}
	table.CreateTable(c)
	quitButton := widget.NewButton("Quit", func() {
		table.Stop = true
		c.Close()
		quitTo()
	})
	roomId := widget.NewLabel("Room Id:")
	copyButton := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		content := roomLabel.Text
		window.Clipboard().SetContent(content)

	})
	hBox := container.New(layout.NewHBoxLayout(), roomId, roomLabel, copyButton, layout.NewSpacer(), quitButton)
	messageBox := container.New(layout.NewVBoxLayout(), hBox, messageLabel, layout.NewSpacer(), resultLabel)
	titleBox := container.New(layout.NewVBoxLayout(), Title, widget.NewSeparator(), messageBox)
	vBox := container.New(layout.NewVBoxLayout(), titleBox, widget.NewSeparator(), table.Grid())
	go table.GameLoop()
	window.SetContent(vBox)
}
