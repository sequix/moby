// +build linux,cgo,!static_build,journald,journald_compat

package journald // import "github.com/sequix/moby/daemon/logger/journald"

// #cgo pkg-config: libsystemd-journal
import "C"
