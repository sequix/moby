package libcontainerd // import "github.com/sequix/moby/libcontainerd"

import (
	"context"

	"github.com/containerd/containerd"
	"github.com/sequix/moby/libcontainerd/local"
	"github.com/sequix/moby/libcontainerd/remote"
	libcontainerdtypes "github.com/sequix/moby/libcontainerd/types"
	"github.com/sequix/moby/pkg/system"
)

// NewClient creates a new libcontainerd client from a containerd client
func NewClient(ctx context.Context, cli *containerd.Client, stateDir, ns string, b libcontainerdtypes.Backend, useShimV2 bool) (libcontainerdtypes.Client, error) {
	if !system.ContainerdRuntimeSupported() {
		// useShimV2 is ignored for windows
		return local.NewClient(ctx, cli, stateDir, ns, b)
	}
	return remote.NewClient(ctx, cli, stateDir, ns, b, useShimV2)
}
