package idtools // import "github.com/sequix/moby/pkg/idtools"

import (
	"os"

	"github.com/sequix/moby/pkg/system"
)

// This is currently a wrapper around MkdirAll, however, since currently
// permissions aren't set through this path, the identity isn't utilized.
// Ownership is handled elsewhere, but in the future could be support here
// too.
func mkdirAs(path string, mode os.FileMode, owner Identity, mkAll, chownExisting bool) error {
	if err := system.MkdirAll(path, mode); err != nil {
		return err
	}
	return nil
}

// CanAccess takes a valid (existing) directory and a uid, gid pair and determines
// if that uid, gid pair has access (execute bit) to the directory
// Windows does not require/support this function, so always return true
func CanAccess(path string, identity Identity) bool {
	return true
}
