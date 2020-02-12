// +build !linux,!windows

package daemon // import "github.com/sequix/moby/daemon"

func secretsSupported() bool {
	return false
}
