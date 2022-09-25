package main

import (
	"tiktaktoe/pkg/page"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var myWindow fyne.Window
var gameStarted bool

func main() {

	myApp := app.New()
	myWindow = myApp.NewWindow("Tik Tak Toe")

	page.MainMenu(myWindow)
	myWindow.ShowAndRun()
}
