package rainbow

import (
	"runtime"
)

// Ver prints the Go version
func Ver() string {
	return runtime.Version()
}
