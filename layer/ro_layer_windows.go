package layer // import "github.com/sequix/moby/layer"

import "github.com/sequix/distribution"

var _ distribution.Describable = &roLayer{}

func (rl *roLayer) Descriptor() distribution.Descriptor {
	return rl.descriptor
}
