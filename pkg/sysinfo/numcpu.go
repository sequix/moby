// +build !linux,!windows

package sysinfo // import "github.com/sequix/moby/pkg/sysinfo"

import (
	"runtime"
)

// NumCPU returns the number of CPUs
func NumCPU() int {
	return runtime.NumCPU()
}
