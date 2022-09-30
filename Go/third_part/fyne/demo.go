package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("hi")

	hi := widget.NewLabel("hi there")
	w.SetContent(container.NewVBox(
		hi,
		widget.NewButton("hi", func() {
			hi.SetText("___:>")
		}),
	))
	w.ShowAndRun()
}
