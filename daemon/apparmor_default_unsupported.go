// +build !linux

package daemon // import "github.com/sequix/moby/daemon"

func ensureDefaultAppArmorProfile() error {
	return nil
}
