// +build !linux,!windows

package daemon // import "github.com/sequix/moby/daemon"

func configsSupported() bool {
	return false
}
