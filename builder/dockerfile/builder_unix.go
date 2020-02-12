// +build !windows

package dockerfile // import "github.com/sequix/moby/builder/dockerfile"

func defaultShellForOS(os string) []string {
	return []string{"/bin/sh", "-c"}
}
