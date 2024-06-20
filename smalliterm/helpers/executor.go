package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"small_projects_in_go/smalliterm/commands"
	"strings"
)

var history []string

func RunCommand(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}

	switch commands.StringToCommand(parts[0]) {
	case commands.ChangeDirectoryCommand:
		commands.ChangeDir(parts)
	case commands.ShowHistoryCommand:
		commands.ShowHistory(history)
	default:
		cmd := exec.Command(parts[0], parts[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error %v\n", err)
		}
	}
	history = append(history, input)
}
