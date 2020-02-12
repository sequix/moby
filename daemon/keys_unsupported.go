// +build !linux

package daemon // import "github.com/sequix/moby/daemon"

// ModifyRootKeyLimit is a noop on unsupported platforms.
func ModifyRootKeyLimit() error {
	return nil
}
