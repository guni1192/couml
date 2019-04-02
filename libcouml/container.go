package libcouml

import (
	"os"
	"os/exec"
	"syscall"
)

// Container -- for base container struct.
// not support expect for Linux
type linuxContainer struct {
	ID string
}

// Container -- for Container Utility
type Container interface {
	// Run -- Container run
	Run(process Process) error
}

// NewContainer -- return Container from linuxContainer
func NewContainer() Container {
	return &linuxContainer{}
}

// Run -- Container Run
func (c *linuxContainer) Run(process Process) error {
	// TODO: prepare somethingCommand  for container init

	if len(os.Args) < 2 {
		// exec itself
		cmd := exec.Command(os.Args[0], "--child")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	cmd := exec.Command(process.Command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUSER | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUTS,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      syscall.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      syscall.Getgid(),
				Size:        1,
			},
		},
	}

	return cmd.Run()
}
