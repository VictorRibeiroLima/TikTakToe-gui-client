package page

import (
	"tiktaktoe/pkg/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MainMenu(window fyne.Window) {

	cR := widget.NewButton("Create Room", func() {
		createRoom(window)
	})
	joinRoom := widget.NewButton("Join Room", func() {
		joinRoom(window)
	})

	hBox := container.New(layout.NewHBoxLayout(), cR, layout.NewSpacer(), joinRoom)
	vBox := container.New(layout.NewVBoxLayout(), components.Title, widget.NewSeparator(), hBox)
	window.SetContent(vBox)
}
