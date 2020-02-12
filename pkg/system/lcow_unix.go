// +build !windows

package system // import "github.com/sequix/moby/pkg/system"

// LCOWSupported returns true if Linux containers on Windows are supported.
func LCOWSupported() bool {
	return false
}
