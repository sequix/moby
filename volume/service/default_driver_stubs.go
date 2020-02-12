// +build !linux,!windows

package service // import "github.com/sequix/moby/volume/service"

import (
	"github.com/sequix/moby/pkg/idtools"
	"github.com/sequix/moby/volume/drivers"
)

func setupDefaultDriver(_ *drivers.Store, _ string, _ idtools.Identity) error { return nil }
