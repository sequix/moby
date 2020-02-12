// +build linux freebsd darwin openbsd

package layer // import "github.com/sequix/moby/layer"

import "github.com/sequix/moby/pkg/stringid"

func (ls *layerStore) mountID(name string) string {
	return stringid.GenerateRandomID()
}
