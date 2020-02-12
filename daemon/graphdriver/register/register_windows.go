package register // import "github.com/sequix/moby/daemon/graphdriver/register"

import (
	// register the windows graph drivers
	_ "github.com/sequix/moby/daemon/graphdriver/lcow"
	_ "github.com/sequix/moby/daemon/graphdriver/windows"
)
