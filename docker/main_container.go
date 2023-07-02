//go:build linux
// +build linux

/*
**
Setup a container-like environment by using Linux namespaces, changing the root directory,
setting up cgroups for resource management, and running a command within this environment.
The run function is responsible for the parent process setup,
while the child function prepares and executes the child process within the container environment.
**
*/
package docker

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

//docker run image <docker> <params>
//go run main_container.go run <docker> <params>

func Main_Linux() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("Bad Command")
	}
}

func run() {
	fmt.Printf("Running %v\n", os.Args[2:])
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	cmd.Run()
}

func child() {
	fmt.Printf("Running %v\n", os.Args[2:])

	cg()

	syscall.Sethostname([]byte("container"))
	syscall.Chroot("/home/fernas/ubuntufs")
	os.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()

	syscall.Unmount("proc", 0)
}

func cg() {
	cgroups := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroups, "pids")
	os.Mkdir(filepath.Join(pids, "fernas"), 0755)
	//:() { : | : & }; : create infinite processes
	ioutil.WriteFile(filepath.Join(pids, "fernas/pids.max"), []byte("20"), 0700)
	// Removes the new cgroup in place after the container exits
	ioutil.WriteFile(filepath.Join(pids, "fernas/notify_on_release"), []byte("1"), 0700)
	ioutil.WriteFile(filepath.Join(pids, "fernas/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700)
}
