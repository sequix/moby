package daemon // import "github.com/sequix/moby/daemon"

import (
	"testing"

	"github.com/sequix/moby/api/types"
	"github.com/sequix/moby/dockerversion"
	"gotest.tools/assert"
)

func TestFillLicense(t *testing.T) {
	v := &types.Info{}
	d := &Daemon{
		root: "/var/lib/docker/",
	}
	d.fillLicense(v)
	assert.Assert(t, v.ProductLicense == dockerversion.DefaultProductLicense)
}
