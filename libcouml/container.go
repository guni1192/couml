package libcouml

import (
	"log"
	"os"
	"path/filepath"
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
	Run(process *Process) error
}

// NewContainer -- return Container from linuxContainer
func NewContainer() Container {
	return &linuxContainer{}
}

// Run -- Container Run
func (c *linuxContainer) Run(process *Process) error {

	if err := syscall.Chroot(process.Cwd); err != nil {
		log.Fatal("chroot failed")
	}
	syscall.Chdir("/")

	return syscall.Exec(process.Args[0], process.Args, process.Env)
}

// PrepareRootfs -- mount file system, change hostname
func PrepareRootfs(config *ContainerConfig) {
	proc := filepath.Join(config.Cwd, "/proc")
	if _, err := os.Stat(proc); os.IsNotExist(err) {
		if err = os.MkdirAll(proc, 0755); err != nil {
			log.Fatalf("mkdir %s failed: %s", proc, err)
		}
	}

	if err := syscall.Mount("proc", proc, "proc", 0, ""); err != nil {
		log.Fatal("cannot mount procfs:", err)
	}
}
