package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/container"
	"fmt"
)

func main() {
	app := app.New()
	window := app.NewWindow("Hello World")
	window.Resize(fyne.NewSize(200, 100))
	window.SetFixedSize(true)

	label := widget.NewLabel("Name:")

	input := widget.NewEntry()
	input.SetPlaceHolder("Your name")
	
	button := widget.NewButton("click me", func() {
		fmt.Println(input.Text)
		input.SetText("")
	})

	grid := container.New(layout.NewVBoxLayout(), label, input, widget.NewLabel(""), button)
	pad := container.New(layout.NewPaddedLayout(), grid)
	
	window.SetContent(pad)
	
	window.ShowAndRun()
}
