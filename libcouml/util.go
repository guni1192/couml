package libcouml

import (
	"os"
	"strings"
)

func setEnv(env []string) {
	for _, e := range env {
		// PATH=/bin -> ["PATH", "/bin"]
		envSet := strings.Split(e, "=")
		if len(envSet) > 1 {
			os.Setenv(envSet[0], envSet[1])
		}
	}
}
