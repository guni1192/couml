package libcouml

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/opencontainers/runtime-spec/specs-go"
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
	command, err := exec.LookPath(process.Args[0])

	if err != nil {
		log.Fatalf("Could not LookPath %s : %s", process.Args[0], err)
	}

	return syscall.Exec(command, process.Args, process.Env)
}

// PrepareRootfs -- mount file system, change hostname
func PrepareRootfs(spec *specs.Spec) {
	proc := filepath.Join(spec.Root.Path, "/proc")
	if _, err := os.Stat(proc); os.IsNotExist(err) {
		if err = os.MkdirAll(proc, 0755); err != nil {
			log.Fatalf("Could not mkdir %s: %s", proc, err)
		}
	}

	if err := syscall.Mount("proc", proc, "proc", 0, ""); err != nil {
		log.Fatal("Could not mount procfs:", err)
	}

	if err := syscall.Chroot(spec.Root.Path); err != nil {
		log.Fatal("Could not chroot", err)
	}
	if err := syscall.Chdir("/"); err != nil {
		log.Fatal("Could not chdir / :", err)
	}
}
