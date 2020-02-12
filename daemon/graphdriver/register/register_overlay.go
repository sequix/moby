// +build !exclude_graphdriver_overlay,linux

package register // import "github.com/sequix/moby/daemon/graphdriver/register"

import (
	// register the overlay graphdriver
	_ "github.com/sequix/moby/daemon/graphdriver/overlay"
)
