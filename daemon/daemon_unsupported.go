// +build !linux,!freebsd,!windows

package daemon // import "github.com/sequix/moby/daemon"
import "github.com/sequix/moby/daemon/config"

const platformSupported = false

func setupResolvConf(config *config.Config) {
}
