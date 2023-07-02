package commandLine

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Main_Terminal() {
	application := app.New()
	window := application.NewWindow("SSH")

	txtGrid := widget.NewTextGrid()
	txtGrid.SetText("This is the terminal")
	window.SetContent(
		container.New(
			layout.NewGridWrapLayout(fyne.NewSize(420, 200)),
			txtGrid,
		),
	)
	window.ShowAndRun()
}
