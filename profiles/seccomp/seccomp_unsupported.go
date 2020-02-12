// +build linux,!seccomp

package seccomp // import "github.com/sequix/moby/profiles/seccomp"

import (
	"github.com/sequix/moby/api/types"
)

// DefaultProfile returns a nil pointer on unsupported systems.
func DefaultProfile() *types.Seccomp {
	return nil
}
