package main

import (
	"os"
	"small_projects_in_go/commandLine"
)

func main() {
	program := os.Args[1]

	switch program {
	case "terminal":
		commandLine.Main_Terminal()
	case "docker":
		/*if runtime.GOOS == "linux" {
			docker.Main_Linux()
		}*/
	}
}
