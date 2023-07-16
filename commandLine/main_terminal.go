package commandLine

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/creack/pty"
	"os"
	"os/exec"
	"time"
)

func Main_Terminal() {
	application := app.New()
	window := application.NewWindow("SSH")

	txtGrid := widget.NewTextGrid()
	txtGrid.SetText("This is the terminal")

	kernelBash(txtGrid, window)

	window.SetContent(
		container.New(
			layout.NewGridWrapLayout(fyne.NewSize(420, 200)),
			txtGrid,
		),
	)
	window.ShowAndRun()
}

func kernelBash(txtGrid *widget.TextGrid, window fyne.Window) {
	command := exec.Command("/bin/bash")
	ptyCommunicator, err := pty.Start(command)
	if err != nil {
		fyne.LogError("Faioed to open pty", err)
		os.Exit(1)
	}

	defer command.Process.Kill()
	processKey(ptyCommunicator, window)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			b := make([]byte, 256)
			_, err := ptyCommunicator.Read(b)
			if err != nil {
				fyne.LogError("Failed to read from pty", err)
			}
			txtGrid.SetText(string(b))
		}
	}()
}

func processKey(ptyCommunicator *os.File, window fyne.Window) {
	onTypedKey := func(key *fyne.KeyEvent) {
		if key.Name == fyne.KeyEnter || key.Name == fyne.KeyReturn {
			ptyCommunicator.Write([]byte{'\r'})
		}
	}

	onTypedRune := func(rune rune) {
		ptyCommunicator.WriteString(string(rune))
	}

	window.Canvas().SetOnTypedKey(onTypedKey)
	window.Canvas().SetOnTypedRune(onTypedRune)
}
