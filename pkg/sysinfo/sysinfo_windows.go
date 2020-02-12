package sysinfo // import "github.com/sequix/moby/pkg/sysinfo"

// New returns an empty SysInfo for windows for now.
func New(quiet bool) *SysInfo {
	sysInfo := &SysInfo{}
	return sysInfo
}
