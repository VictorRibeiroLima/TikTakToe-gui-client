package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

var roomLabel = widget.NewLabel("")
var messageLabel = widget.NewLabel("")
var resultLabel = widget.NewLabel("")
var Title = &canvas.Text{
	Text:  "Tik Tak Toe",
	Color: color.White,
	TextStyle: fyne.TextStyle{
		Bold: true,
	},
	Alignment: fyne.TextAlignCenter,
	TextSize:  24,
}
