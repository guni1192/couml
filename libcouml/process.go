package libcouml

import (
	"github.com/opencontainers/runtime-spec/specs-go"
)

// Process -- have process state
type Process struct {
	Cwd     string // Current working directory
	Env     []string
	Args    []string
}

// NewProcess --  generate new process config.
func NewProcess(p *specs.Process) *Process {
	return &Process{
		Cwd:  p.Cwd,
		Env:  p.Env,
		Args: p.Args,
	}
}
