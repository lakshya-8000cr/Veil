package overlay

import (
	"os"
	"strings"
)

func IsMounted(target string) bool {  // this is for the inspect command
	data, err := os.ReadFile("/proc/mounts")
	if err != nil {
		return false
	}

	return strings.Contains(string(data), target)
}