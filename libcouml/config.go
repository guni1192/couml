package libcouml

// ContainerConfig -- container initial config.
// use loading config.json
type ContainerConfig struct {
	Args []string `json:"args"`
	Env  []string `json:"env"`
	Cwd  string   `json:"cwd"`
}
