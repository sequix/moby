package daemon // import "github.com/sequix/moby/daemon"

import (
	"github.com/sequix/moby/api/types/container"
	libcontainerdtypes "github.com/sequix/moby/libcontainerd/types"
)

func toContainerdResources(resources container.Resources) *libcontainerdtypes.Resources {
	// We don't support update, so do nothing
	return nil
}
