package libcouml

import (
	"encoding/json"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	"io/ioutil"
	"log"
)

// LoadConfig -- Load spec info from configPath
func LoadConfig(configPath string) *specs.Spec {
	raw, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Could not read %s: %s", configPath, err)
	}

	s := specs.Spec{}

	if err = json.Unmarshal(raw, &s); err != nil {
		log.Fatal("json Unmarshal Error: ", err)
	}

	return &s
}
