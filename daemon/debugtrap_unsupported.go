// +build !linux,!darwin,!freebsd,!windows

package daemon // import "github.com/sequix/moby/daemon"

func (d *Daemon) setupDumpStackTrap(_ string) {
	return
}
