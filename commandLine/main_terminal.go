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

	os.Setenv("TERM", "dumb")

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
	command := exec.Command("/bin/sh")
	ptyCommunicator, err := pty.Start(command)
	if err != nil {
		fyne.LogError("Faioed to open pty", err)
		os.Exit(1)
	}

	_, err = ptyCommunicator.WriteString("ls\r")
	if err != nil {
		fyne.LogError(err.Error(), err)
	}

	b := make([]byte, 1024)
	_, err = ptyCommunicator.Read(b)

	if err != nil {
		fyne.LogError(err.Error(), err)
	}

	txtGrid.SetText(string(b))

	processKey(ptyCommunicator, window)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			b = make([]byte, 1024)
			n, err := ptyCommunicator.Read(b)
			if n > 0 && err != nil {
				fyne.LogError(err.Error(), err)
			}
			txtGrid.SetText(string(b))
		}
	}()

	defer command.Process.Kill()

}

// processKey callbacks that handles special keypresses and characters keypresses
func processKey(ptyCommunicator *os.File, window fyne.Window) {
	onTypedKey := func(key *fyne.KeyEvent) {
		if key.Name == fyne.KeyEnter || key.Name == fyne.KeyReturn {
			_, _ = ptyCommunicator.Write([]byte{'\r'})
		}
	}

	onTypedRune := func(rune rune) {
		_, _ = ptyCommunicator.WriteString(string(rune))
	}

	window.Canvas().SetOnTypedKey(onTypedKey)
	window.Canvas().SetOnTypedRune(onTypedRune)
}
