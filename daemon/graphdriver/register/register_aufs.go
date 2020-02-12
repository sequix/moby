// +build !exclude_graphdriver_aufs,linux

package register // import "github.com/sequix/moby/daemon/graphdriver/register"

import (
	// register the aufs graphdriver
	_ "github.com/sequix/moby/daemon/graphdriver/aufs"
)
