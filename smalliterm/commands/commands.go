package commands

import (
	"fmt"
	"os"
)

// ChangeDir changes the current working directory
func ChangeDir(parts []string) {
	if len(parts) < 2 {
		_, _ = fmt.Fprintf(os.Stderr, "cd needs an argument\n")
		return
	}
	err := os.Chdir(parts[1]) // changes directory
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error %v\n", err)
	}
}

// ShowHistory prints the history of commands
func ShowHistory(history []string) {
	for i, cmd := range history {
		fmt.Printf("%d: %s\n", i, cmd)
	}
}

func StringToCommand(s string) Command {
	switch s {
	case "cd":
		return ChangeDirectoryCommand
	case "history":
		return ShowHistoryCommand
	default:
		return "" // or some default value
	}
}
