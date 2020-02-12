package image

import (
	"context"
	"testing"

	"github.com/sequix/moby/api/types"
	"github.com/sequix/moby/api/types/versions"
	"github.com/sequix/moby/errdefs"
	"gotest.tools/assert"
	"gotest.tools/skip"
)

func TestImagePullPlatformInvalid(t *testing.T) {
	skip.If(t, versions.LessThan(testEnv.DaemonAPIVersion(), "1.40"), "experimental in older versions")
	defer setupTest(t)()
	client := testEnv.APIClient()
	ctx := context.Background()

	_, err := client.ImagePull(ctx, "docker.io/library/hello-world:latest", types.ImagePullOptions{Platform: "foobar"})
	assert.Assert(t, err != nil)
	assert.ErrorContains(t, err, "unknown operating system or architecture")
	assert.Assert(t, errdefs.IsInvalidParameter(err))
}
