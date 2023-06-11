package main

import (
	"os"
	"os/exec"
	"syscall"
)

//docker run image <main> <params>
//go run main_linux.go run <main> <params>

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("Bad Command")
	}
}

func run() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	cmd.Run()
}
