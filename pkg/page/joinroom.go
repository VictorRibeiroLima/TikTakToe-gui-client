package page

import (
	"tiktaktoe/pkg/components"
	"tiktaktoe/pkg/internal"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func joinRoom(window fyne.Window) {
	var id string
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter Room id...")
	hBox := container.NewVBox(input, layout.NewSpacer(), widget.NewButton("Save", func() {
		id = input.Text
		c, _, err := internal.Connect("/ws/join-room/" + id)
		if err != nil {
			MainMenu(window)
		}
		components.CreateLobby(window, c, func() {
			MainMenu(window)
		})
	}))
	vBox := container.New(layout.NewVBoxLayout(), components.Title)
	c := fyne.NewContainerWithLayout(layout.NewBorderLayout(vBox, hBox, nil, nil), vBox, hBox)
	window.SetContent(c)
}
