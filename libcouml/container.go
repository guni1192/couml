package libcouml

import "fmt"

// Container -- for base container struct.
// not support expect for Linux
type linuxContainer struct {
	ID string
}

// Container -- for Container Utility
type Container interface {
	// Run -- Container run
	Run() error
}

// NewContainer -- return Container from linuxContainer
func NewContainer() Container {
	return &linuxContainer{}
}

// Run -- Container Run
func (c *linuxContainer) Run() error {
	fmt.Println("Container Run!")
	return nil
}
