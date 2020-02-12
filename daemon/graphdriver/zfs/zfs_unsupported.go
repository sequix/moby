// +build !linux,!freebsd

package zfs // import "github.com/sequix/moby/daemon/graphdriver/zfs"

func checkRootdirFs(rootdir string) error {
	return nil
}

func getMountpoint(id string) string {
	return id
}
