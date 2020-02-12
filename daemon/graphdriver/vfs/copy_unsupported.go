// +build !linux

package vfs // import "github.com/sequix/moby/daemon/graphdriver/vfs"

import "github.com/sequix/moby/pkg/chrootarchive"

func dirCopy(srcDir, dstDir string) error {
	return chrootarchive.NewArchiver(nil).CopyWithTar(srcDir, dstDir)
}
