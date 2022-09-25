package page

import (
	"tiktaktoe/pkg/components"
	"tiktaktoe/pkg/internal"

	"fyne.io/fyne/v2"
)

func createRoom(window fyne.Window) {
	c, _, err := internal.Connect("/ws/create-room/")
	if err != nil {
		MainMenu(window)
	}
	components.CreateLobby(window, c, func() {
		MainMenu(window)
	})
}
