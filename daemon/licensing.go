package daemon // import "github.com/sequix/moby/daemon"

import (
	"github.com/sequix/moby/api/types"
	"github.com/sequix/moby/dockerversion"
)

func (daemon *Daemon) fillLicense(v *types.Info) {
	v.ProductLicense = dockerversion.DefaultProductLicense
}
