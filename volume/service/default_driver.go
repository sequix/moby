// +build linux windows

package service // import "github.com/sequix/moby/volume/service"
import (
	"github.com/sequix/moby/pkg/idtools"
	"github.com/sequix/moby/volume"
	"github.com/sequix/moby/volume/drivers"
	"github.com/sequix/moby/volume/local"
	"github.com/pkg/errors"
)

func setupDefaultDriver(store *drivers.Store, root string, rootIDs idtools.Identity) error {
	d, err := local.New(root, rootIDs)
	if err != nil {
		return errors.Wrap(err, "error setting up default driver")
	}
	if !store.Register(d, volume.DefaultDriverName) {
		return errors.New("local volume driver could not be registered")
	}
	return nil
}
