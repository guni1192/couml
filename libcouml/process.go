package libcouml

// Process -- have process state
type Process struct {
	Command string
	Cwd     string // Current working directory
}

// NewProcess --  generate new process config.
// TODO: replace LoadConfig from config.json
func NewProcess(cmd string, workDir string) *Process {
	return &Process{
		Command: cmd,
		Cwd:     workDir,
	}
}
