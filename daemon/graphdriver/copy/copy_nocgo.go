// +build linux,!cgo

package copy // import "github.com/sequix/moby/daemon/graphdriver/copy"

import (
	"os"

	"golang.org/x/sys/unix"
)

func fiClone(srcFile, dstFile *os.File) error {
	return unix.ENOSYS
}
