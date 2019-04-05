package libcouml

import (
	"github.com/opencontainers/runtime-spec/specs-go"
)

// Process -- have process state
type Process struct {
	Command string // TODO: remove
	Cwd     string // Current working directory
	Env []string
	Args []string
}

// NewProcess --  generate new process config.
// TODO: replace LoadConfig from config.json
func NewProcess(args []string, workDir string) *Process {
	return &Process{
		Args: args,
		Cwd:     workDir,
	}
}

// newProcess -- TODO: instead of NewProcess
func newProcess(p *specs.Process) *Process{
	return &Process {
		Cwd: p.Cwd,
		Env: p.Env,
		Args: p.Args,
	}
}
