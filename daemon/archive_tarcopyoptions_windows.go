package daemon // import "github.com/sequix/moby/daemon"

import (
	"github.com/sequix/moby/container"
	"github.com/sequix/moby/pkg/archive"
)

func (daemon *Daemon) tarCopyOptions(container *container.Container, noOverwriteDirNonDir bool) (*archive.TarOptions, error) {
	return daemon.defaultTarCopyOptions(noOverwriteDirNonDir), nil
}
