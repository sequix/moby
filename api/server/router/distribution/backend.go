package distribution // import "github.com/sequix/moby/api/server/router/distribution"

import (
	"context"

	"github.com/sequix/distribution"
	"github.com/sequix/distribution/reference"
	"github.com/sequix/moby/api/types"
)

// Backend is all the methods that need to be implemented
// to provide image specific functionality.
type Backend interface {
	GetRepository(context.Context, reference.Named, *types.AuthConfig) (distribution.Repository, bool, error)
}
