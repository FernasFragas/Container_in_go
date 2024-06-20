package helpers

import (
	"github.com/c-bata/go-prompt"
	"os"
	"path/filepath"
	"strings"
)

func GetPrompt() string {
	cwd, err := os.Getwd()
	if err != nil {
		return ">"
	}
	return filepath.Base(cwd) + " > "
}

func ChangeLivePrefix() (string, bool) {
	return GetPrompt(), true
}

func Executor(input string) {
	input = strings.TrimSpace(input)
	if input == "" {
		return
	}
	if input == "exit" {
		os.Exit(0)
	}
	RunCommand(input)
}

func Completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: "exit", Description: "exit the shell"},
		{Text: "ls", Description: "List directory contents"},
		{Text: "cd", Description: "Change the current working directory"},
		{Text: "pwd", Description: "Print the current working directory"},
		{Text: "echo", Description: "Prints the arguments to the standard output"},
		{Text: "cat", Description: "Concatenate files and print on the standard output"},
		{Text: "wc", Description: "Print newline, word, and byte counts for each file"},
		{Text: "grep", Description: "Print lines that match patterns"},
		{Text: "history", Description: "Print the history of commands"},
		{Text: "clear", Description: "Clear the screen"},
	}
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}
