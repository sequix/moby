package vfs // import "github.com/sequix/moby/daemon/graphdriver/vfs"

import "github.com/sequix/moby/daemon/graphdriver/copy"

func dirCopy(srcDir, dstDir string) error {
	return copy.DirCopy(srcDir, dstDir, copy.Content, false)
}
