package main

import (
	"os"
	"small_projects_in_go/smalliterm"
)

func main() {
	program := os.Args[1]

	switch program {
	case "iterm":
		smalliterm.Main_Iterm()
	}
}
