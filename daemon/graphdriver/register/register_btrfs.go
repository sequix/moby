// +build !exclude_graphdriver_btrfs,linux

package register // import "github.com/sequix/moby/daemon/graphdriver/register"

import (
	// register the btrfs graphdriver
	_ "github.com/sequix/moby/daemon/graphdriver/btrfs"
)
